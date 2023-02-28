package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
			fmt.Println("Creating file " + getContent.Content + ".json")
			// File doesn't exist, create it
			jsonFile, err = os.Create(dir)
			if err != nil {
				PrintErr("Error creating file " + getContent.Content + ".json" + err.Error())
				return
			}
			// Write the JSON string to the file
			jsonString := "{\"type\":\"doc\",\"content\":[{\"type\":\"paragraph\",\"content\":[{\"type\":\"text\",\"text\":\"Example Text\"}]}]}"
			_, err = jsonFile.WriteString(jsonString)
			if err != nil {
				PrintErr("Error writing to file " + getContent.Content + ".json" + err.Error())
				return
			}
			PrintSuccess(dir + " created")
			byteValue = []byte(jsonString)
		} else {
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
		err = json.Unmarshal(byteValue, &datas)
		if err != nil {
			PrintErr("Error parsing JSON:" + err.Error())
			return
		}
		fmt.Fprintln(w, string(byteValue))
		PrintSuccess("read from file " + dir + " successful")
	}
}
func HandlerSaveContent(StoreDir string, dirs map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// If request is preflight, return
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println()
		var content Content
		// Read json data
		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		err := json.Unmarshal((b), &content)
		if err != nil {
			PrintErr("Error parsing JSON:" + err.Error())
			return
		}
		fmt.Println("saving to", content.Content, "on", content.Page)

		// Store json data
		dir := StoreDir + "/" + content.Page + "/" + content.Content + ".json"
		jsonFile, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			PrintErr("Error open " + dir + ":" + err.Error())
			return
		}
		// Remove previous data
		jsonFile.Truncate(0)
		_, err = jsonFile.Seek(0, 0)
		if err != nil {
			PrintErr("Error seeking to beginning of file:" + err.Error())
			return
		}
		_, err = jsonFile.Write([]byte(string(content.Contentjson)))
		if err != nil {
			fmt.Println(err)
		}
		jsonFile.Close()

		// Incert html data to file
		// Open the file using the path stored in the dirs map
		file, err := os.OpenFile(dirs[content.Page+"?"+content.Content], os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			PrintErr("Error opening file" + dir + ":" + err.Error())
			return
		}
		defer file.Close()
		// Read the file contents
		b, err = io.ReadAll(file)
		if err != nil {
			PrintErr("Error reading file:" + err.Error())
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
			tagAfter = "<!-- WikiBreeze " + content.Content + " end-->\n" + string(b[startPattern[1]+1:])
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
			return
		}
		// Write the modified contents back to the file
		_, err = file.Write(newContents)
		if err != nil {
			PrintErr("Error writing to file:" + err.Error())
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
			return
		}
		fmt.Println()

		fmt.Fprint(w, string(RenderConfigString))
	}
}
