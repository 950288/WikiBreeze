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
	dir := "D:\\github\\web\\src\\pages"
	re := regexp.MustCompile(`<!--\s*iGEMGotool\s*(?P<name>\S+)\s*start-->`)

	// Create map to store data
	data := make(map[string][]string)

	//scan *.vue in dir
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".vue" {
			// Read file contents
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Extract "name" value from file contents
			matches := re.FindAllStringSubmatch(string(b) ,-1)

			// Store file name and "name" values in map
			fileName := strings.TrimSuffix(filepath.Base(path), ".vue")
			nameValues := make([]string, len(matches))


			for i , match := range matches {
				nameValues[i] = match[1]
			}
			data[fileName] = nameValues
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fmt.Fprintf(w, string(jsonData))
		fmt.Println("some one visit")
	})
	http.HandleFunc("/getnode", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
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

		fmt.Println("editing ",node," on",page)

		//read json data and create it if it doesn't exist
		dir := "./data/"+page+"/"+node+".json"
		err = os.MkdirAll("./data/"+page, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
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
		fmt.Println("reading:\t", string(byteValue))
		err = json.Unmarshal(byteValue, &datas)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		fmt.Fprintf(w, string(byteValue))



	})
	http.Handle("/", http.FileServer(http.Dir("../dist")))
	http.ListenAndServe(":8082", nil)

}

// `<!-- iGEMGotool {{name}} start-->`
