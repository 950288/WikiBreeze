package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
)

type GetContent struct {
	Page    string
	Content string
}
type Content struct {
	Page        string
	Content     string
	Contentjson json.RawMessage
	Contenthtml string
}

func HandlerFetchContentList(getContentByte []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//If request is Preflight, return
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println()
		fmt.Println("content list fetched")

		fmt.Fprint(w, string(getContentByte))
	}
}
func HandlerGetContent(StoreDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//If request is Preflight, return
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println()

		var byteValue []byte
		var getContent GetContent
		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		err := json.Unmarshal([]byte(string(b)), &getContent)
		if err != nil {
			PrintErr(err.Error())
			return
		}

		fmt.Println("editing", getContent.Content, "on", getContent.Page)
		//read json data and create it if it doesn't exist
		dir := StoreDir + "/" + getContent.Page + "/" + getContent.Content + ".json"
		err = os.MkdirAll(StoreDir+"/"+getContent.Page, os.ModePerm)
		if err != nil {
			PrintErr("Error creating directory:" + err.Error())
			return
		}
		//read or create json file
		jsonFile, err := os.Open(dir)
		if err != nil {
			fmt.Println(dir + " doesn't exist , reading from default")
			// fmt.Println("Creating file " + getContent.Content + ".json")
			// File doesn't exist, create it
			// jsonFile, err = os.Create(dir)
			// if err != nil {
			// 	PrintErr("Error creating file " + getContent.Content + ".json" + err.Error())
			// 	return
			// }

			var jsonString string
			fmt.Println(getContent.Content, getContent.Page)
			// fmt.Println(TestPageJson)
			if getContent.Content == "content" && getContent.Page == "testPage" {
				fin, err := os.Open("./testContent.json")
				if err != nil {
					log.Fatal(fmt.Errorf("error open to testContent.json: %w", err))
				}
				defer fin.Close()
				byteValue, err = io.ReadAll(fin)
				if err != nil {
					log.Fatal(fmt.Errorf("error read to testContent.json: %w", err))
				}
				jsonString = string(byteValue)
			} else {
				jsonString = DefaultRenderJson
			}
			// Write the JSON string to the file
			// _, err = jsonFile.WriteString(jsonString)
			// if err != nil {
			// 	PrintErr("Error writing to file " + getContent.Content + ".json" + err.Error())
			// 	return
			// }
			// PrintSuccess(dir + " created")
			byteValue = []byte(jsonString)
		} else {
			PrintSuccess("read from file " + dir)
			defer jsonFile.Close()
			byteValue, err = io.ReadAll(jsonFile)
			if err != nil {
				PrintErr("Error reading file " + getContent.Content + ".json " + err.Error())
				return
			}
		}
		jsonFile.Close()

		// Parse the JSON data
		var datas map[string]interface{}
		if string(byteValue) == "" {
			byteValue = []byte("{}")
		}
		err = json.Unmarshal(byteValue, &datas)
		if err != nil {
			PrintErr("Error parsing " + dir + ":" + err.Error())
			return
		}
		// fmt.Println(string(byteValue))
		fmt.Fprintln(w, string(byteValue))
		PrintSuccess("fetch successful")
	}
}
func HandlerSaveContent(StoreDir string, dirs map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// If request is preflight, return
		if r.Method == "OPTIONS" {
			r.Body.Close()
			return
		}

		// Set a max request size
		// r.Body = http.MaxBytesReader(w, r.Body, maxRequestSize)

		var count = 0
		var content Content
		decoder := json.NewDecoder(r.Body)
		for {

			var c Content
			err := decoder.Decode(&c)
			if err == io.EOF {
				break
			} else if err != nil {
				PrintErr("Error parsing JSON: " + err.Error())
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			count += 1
			// PrintSuccess(c.Content)
			// fmt.Println(c.Contenthtml)
			content.Content += c.Content
			content.Page += c.Page
			// Merge JSON
			if content.Contentjson == nil {
				content.Contentjson = c.Contentjson
			} else if c.Contentjson != nil {
				content.Contentjson = append(content.Contentjson, ',')
				content.Contentjson = append(content.Contentjson, c.Contentjson...)
			}

			// Concatenate HTML
			if c.Contenthtml != "" {
				content.Contenthtml += c.Contenthtml
			}
		}

		// Use content...
		fmt.Println("\nSaving to", content.Content, "on", content.Page, "(", count, ")")

		r.Body.Close()

		// Open file and write JSON
		dir := StoreDir + "/" + content.Page + "/" + content.Content + ".json"
		jsonFile, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			PrintErr("Error opening " + dir + ": " + err.Error())
			fmt.Fprint(w, "{\"success\": \""+"Error opening "+dir+": "+err.Error()+"\"}")
			return
		}
		// Remove previous data
		jsonFile.Truncate(0)
		_, err = jsonFile.Seek(0, 0)
		if err != nil {
			PrintErr("Error seeking to beginning of file:" + err.Error())
			fmt.Fprint(w, "{\"success\": \""+"Error seeking to beginning of file:"+err.Error()+"\"}")
			return
		}
		// Write the modified contents back to the file
		// MarshalIndent
		contentString := string(content.Contentjson)
		contentString, err = ParseJson("./src/config.json", contentString)
		if err != nil {
			fmt.Println("MarshalIndent failed:", err.Error())
		}
		_, err = jsonFile.Write([]byte(contentString))

		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "{\"success\": \"false\"}")

		}
		jsonFile.Close()

		// Incert html data to file
		// Open the file using the path stored in the dirs map
		file, err := os.OpenFile(dirs[content.Page+"?"+content.Content], os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(dirs[content.Page+"?"+content.Content])
			PrintErr("Error opening file" + dir + ":" + err.Error())
			fmt.Fprint(w, "{\"success\": \"false\"}")
			return
		}
		defer file.Close()
		// Read the file contents
		b, err := io.ReadAll(file)
		if err != nil {
			PrintErr("Error reading file:" + err.Error())
			fmt.Fprint(w, "{\"success\": \"false\"}")

			return
		}
		// Use regular expressions to find the start and end tags
		startTag := regexp.MustCompile(`<!--\s*WikiBreeze\s*` + content.Content + `\s*start-->`)
		endTag := regexp.MustCompile(`<!--\s*WikiBreeze\s*` + content.Content + `\s*end-->`)
		// Replace the contents between the tags with data.Contenthtml
		startPattern := startTag.FindIndex(b)
		tagBefore := string(b[0:startPattern[1]])
		endPattern := endTag.FindIndex(b)
		var tagAfter string
		if len(endPattern) == 0 {
			// If the end tag is not found, append it to the end of the file
			tagAfter = "<!-- WikiBreeze " + content.Content + " end-->\n" + string(b[startPattern[1]:])
		} else {
			// If the end tag is found
			tagAfter = string(b[endPattern[0]:])
		}
		// connect tagBefore, the content data, and tagAfter
		newContents := []byte(tagBefore + content.Contenthtml + tagAfter)
		// remove the old contents
		file.Truncate(0)
		_, err = file.Seek(0, 0)
		if err != nil {
			PrintErr("Error seeking to beginning of file:" + err.Error())
			fmt.Fprint(w, "{\"success\": \"false\"}")
			return
		}
		// Write the modified contents back to the file
		_, err = file.Write(newContents)
		if err != nil {
			PrintErr("Error writing to file:" + err.Error())
			fmt.Fprint(w, "{\"success\": \"false\"}")
			return
		}
		fmt.Fprint(w, "{\"success\": \"true\"}")
		PrintSuccess("saved " + dir + " successfully")

	}
}
func HandlergetRenderconfig(RenderConfigString string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//If request is Preflight, return
		if r.Method == "OPTIONS" {
			r.Body.Close()
			return
		}
		fmt.Println()

		fmt.Fprint(w, string(RenderConfigString))
	}
}
func HandlerUploadImage(cookie *cookiejar.Jar, requestUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//If request is Preflight, return
		if r.Method == "OPTIONS" || r.ContentLength <= 2 {
			r.Body.Close()
			fmt.Fprint(w, "{\"enabled\": \"true\"}")
			return
		}

		fmt.Println("\nunloading files")
		var b []byte
		r.Body.Read(b)

		// reading request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			PrintErr(err.Error())
			fmt.Print(r)
			return
		}

		// Build a forward request
		req, err := http.NewRequest(r.Method, requestUrl, bytes.NewBuffer(body))
		if err != nil {
			PrintErr(err.Error())
			fmt.Print(r)
			return
		}
		req.Header = r.Header

		client := &http.Client{
			CheckRedirect: CheckRedirect, //disable redirect
			Jar:           cookie,
		}

		// Synchronous forward request
		res, err := client.Do(req)
		if err != nil {
			PrintErr(err.Error())
			// fmt.Print(req)
			return
		}
		defer res.Body.Close()

		if res.StatusCode == 201 {
			PrintSuccess("upload pictue successful !")
		} else {
			PrintErr("upload pictue failed :" + res.Status)
		}

		// write the response
		_, err = io.Copy(w, res.Body)
		if err != nil {
			PrintErr(err.Error())
			return
		}

		// set response headers
		for k, v := range res.Header {
			w.Header()[k] = v
		}
	}
}
