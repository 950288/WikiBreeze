package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"os"
	"regexp"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
	"net"
)

// TreeNode represents a node in the tree
type TreeNode struct {
	Subtree []*pages
}
type pages struct {
	name string
	node []*TreeNode
}
func main() {
	scanDir := "..\\"
	storeDir := "..\\iGEMGotoolData"
	reConfig := regexp.MustCompile(`<!--\s*iGEMGotool\s*(?P<name>\S+)\s*start-->`)
	fileTypes := []string{".html",".vue"}
	port := 8080

	configFile, err := os.Open("./config.json")
	if err != nil {
		//create config.json if it doesn't exist
		configFile, err = os.Create("./config.json")
		if err != nil {
			fmt.Println("Error creating config.json", err)
		} else {
			// Write the JSON string to the file
			jsonString := 
`{
	// Directory containing the page to be modified (e.g. "D:\\github\\web\\src\\pages")
	"ScanDirectory": "..\\",

	// Directory to store the edited page (e.g. "D:\\github\\web\\src\\iGEMGotoolData")
	"StoreDirectory": "..\\iGEMGotoolData",

	//Port to be used
	"Port": 8080,

	//the tag to be scan and incert content (e.g. "iGEMGotool"),
	//which be automatically converted to <!-- iGEMGotool {{name}} start-->
	"incert tag":"iGEMGotool",

	
	//file type to be scan (e.g. ".html")
	"file type":[".html",".vue"]
}`
			err = ioutil.WriteFile("./config.json", []byte(jsonString), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}	
			fmt.Println("created config.json")
			defer configFile.Close()
		}
	} else {
		configByteValue, err := ioutil.ReadAll(configFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Parse config.json
		// Use a regular expression to remove comments from the JSON string
		configJsonString := regexp.MustCompile(`(?m)^\s*//.*$|(?m)^\s*/\*[\s\S]*?\*/`).ReplaceAllString(string(configByteValue), "")
		type config struct {
			ScanDirectory string `json:"ScanDirectory"`
			StoreDirectory string `json:"StoreDirectory"`
			Port int `json:"Port"`
			InsertTag string `json:"incert tag"`
			FileTypes []string `json:"file type"`
		}
		var configData config
		// fmt.Println("Parseing:\t", configJsonString)
		err = json.Unmarshal([]byte(configJsonString), &configData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		
		scanDir = configData.ScanDirectory
		storeDir = configData.StoreDirectory
		fileTypes = configData.FileTypes
		port = int(configData.Port)
		reConfig = regexp.MustCompile(`<!--\s*`+ configData.InsertTag+`\s*(?P<name>\S+)\s*start-->`)

		fmt.Println("read config.json successful !")


	}

	// Create map to store data
	data := make(map[string][]string)
	//scan specified type of files in the directory
	filepath.Walk(scanDir, func(path string, info os.FileInfo, err error) error {
		// fileTypes = fileTypes.([]interface{})
		for _, fileType := range fileTypes {	
			if filepath.Ext(path) == fileType {
				// Read file contents
				b, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				// Extract "incert tag name" value from file contents
				matches := reConfig.FindAllStringSubmatch(string(b) ,-1)
				// Store file "incert tag name" values in map
				fileName := strings.TrimSuffix(filepath.Base(path), fileType)
				nameValues := make([]string, len(matches))

				if len(matches) > 0 {
					for i , match := range matches {
						nameValues[i] = match[1]
					}
					data[fileName] = nameValues
				}

			}
		}
		return nil
	})

	// Serialize map to JSON string
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(jsonData) +"\n")
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fmt.Fprintf(w, string(jsonData))
		fmt.Println("some one visit homepage and fetched edit list")
	})
	http.HandleFunc("/getnode", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		type Data struct {
			Page  string
			Node string
		}
		var data Data

		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		err := json.Unmarshal([]byte(string(b)), &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		page := data.Page
		node := data.Node

		fmt.Println("editing",node,"on",page)

		//read json data and create it if it doesn't exist
		dir := storeDir+"/"+page+"/"+node+".json" 
		err = os.MkdirAll(storeDir+"/"+page, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		//read or create json file
		jsonFile, err := os.Open(dir)
		if err != nil {
			fmt.Println("File doesn't exist, creating it...")
			// File doesn't exist, create it
			jsonFile, err = os.Create(dir)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			// Write the JSON string to the file
			jsonString := "{\"type\":\"doc\",\"content\":[{\"type\":\"paragraph\",\"content\":[{\"type\":\"text\",\"text\":\"Example Text\"}]}]}"
			err = ioutil.WriteFile(dir, []byte(jsonString), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}	
			fmt.Println(dir+" created")
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Parse the JSON data
		var datas map[string]interface{}
		fmt.Printf("reading from file %s\n", dir)
		err = json.Unmarshal(byteValue, &datas)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		fmt.Fprintf(w, string(byteValue))
	})
	http.HandleFunc("/savenode", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		type Data struct {
			Page  string
			Node string
			Content json.RawMessage
		}
		var data Data

		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		err := json.Unmarshal([]byte(string(b)), &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		page := data.Page
		node := data.Node
		content := data.Content

		//Store json data
		dir := storeDir+"/"+page+"/"+node+".json" 
		jsonFile, err := os.Open(dir)
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}
		defer jsonFile.Close()
		fmt.Printf("saving to file %s\n", dir)

		err = ioutil.WriteFile(dir, []byte(string(content)), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("saved %s successful\n", dir)
		fmt.Fprintf(w, "success")
	})
	//check if running in production mode
	stat , err := os.Stat("./index.html")
	if stat != nil {
		// fmt.Println("Running in production mode")
		http.Handle("/", http.FileServer(http.Dir("./")))
	} else {
		// fmt.Println("Running in development mode")
		http.Handle("/", http.FileServer(http.Dir("../dist")))
	}
	fmt.Println("Server started on port", strconv.Itoa(port))
	fmt.Println("Local:\t\t http://127.0.0.1:"+strconv.Itoa(port)+"/")
	//get local ip
	if ip := getOutboundIP(); ip != nil {
		fmt.Println("Network:\t http://"+ip.String()+":"+strconv.Itoa(port)+"/")
	}
	
	err = http.ListenAndServe(":"+ strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	} 
}

// getOutboundIP gets the preferred outbound ip of this machine
func getOutboundIP () net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Error getting local IP")
		return nil
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
} 

// `<!-- iGEMGotool {{name}} start-->`
