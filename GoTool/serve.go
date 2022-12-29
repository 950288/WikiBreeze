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
			page  string 	`json:"page"`
			node string		`json:"node"`
		}
		var data Data

		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		fmt.Println(b)
		fmt.Println(string(b))
		err := json.Unmarshal([]byte(string(b)), &data)
		fmt.Println(err)
		fmt.Println(data)
		page := r.URL.Query().Get("page")
		node := r.URL.Query().Get("node")

		// fmt.Fprintf(w, string(jsonData))
		fmt.Println("edit ",node," on",page)

		fmt.Fprintf(w, "Civilized Language")

	})
	http.Handle("/", http.FileServer(http.Dir("../dist")))
	http.ListenAndServe(":8082", nil)

}

// `<!-- iGEM-ToolBox:WIKI {{name}} start-->`
