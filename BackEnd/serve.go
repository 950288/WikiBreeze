package main

import (
	"fmt"
	// "net/http"
	// "path/filepath"
	// "os"
	// "regexp"
	// "io/ioutil"
	// "strings"
	// "encoding/json"
	// "strconv"
	// "net"

	"serve/utils"
)

func main(){
	fff()
	fmt.Println(utils.Config)
}

// func main() {
// 	fmt.Println(utils.config)
// 	utils.fff()
// 	//Initialize basic 
// 	scanDir := "..\\"
// 	storeDir := ".\\WikiData"
// 	port := 8080
// 	PortConfig := "auto"
// 	// reConfig := regexp.MustCompile(`<!--\s*WikiBreeze\s*(?P<name>\S+)\s*start-->`)
// 	tagName := "WikiBreeze"
// 	fileTypes := []string{".html",".vue"}
// 	configFile, err := os.Open("./config.json")
// 	if err != nil {
// 		//create config.json if it doesn't exist
// 		configFile, err = os.Create("./config.json")
// 		if err != nil {
// 			printErr("Error creating config.json" + err.Error())
// 		} else {
// 			// Write the JSON string to the file
// 			jsonString := 
// `{
// 	// Directory containing the page to be modified (e.g. "D:\\github\\web\\src\\pages")
// 	"scanDirectory": "..\\",

// 	// Directory to store the edited page (e.g. "D:\\github\\web\\src\\WikiData")
// 	"storeDirectory": ".\\WikiData",

// 	//Port to be used (e.g. "8080" or "auto")
// 	"port": "auto",

// 	//the tag to be scan and incert content (e.g. "WikiBreeze"),
// 	//which be automatically converted to <!-- WikiBreeze {{name}} start-->
// 	"incertTag":"WikiBreeze",
	
// 	//file type to be scan (e.g. [".html",....])
// 	"fileType":[".html",".vue"]
// }`
// 			err = ioutil.WriteFile("./config.json", []byte(jsonString), 0644)
// 			if err != nil {
// 				printErr("Error writing to file:" + err.Error())
// 				return
// 			}	
// 			fmt.Println("created config.json")
// 			defer configFile.Close()
// 		}
// 	} else {
// 		configByteValue, err := ioutil.ReadAll(configFile)
// 		if err != nil {
// 			printErr("Error reading file:"+ err.Error())
// 			return
// 		}
// 		// Parse config.json
// 		// Use a regular expression to remove comments from the JSON string
// 		configJsonString := regexp.MustCompile(`(?m)^\s*//.*$|(?m)^\s*/\*[\s\S]*?\*/`).ReplaceAllString(string(configByteValue), "")

// 		var configData config
// 		// fmt.Println("Parseing:\t", configJsonString)
// 		err = json.Unmarshal([]byte(configJsonString), &configData)
// 		if err != nil {
// 			printErr("Error parsing JSON:"+ err.Error())
// 			return
// 		}
// 		scanDir = configData.ScanDirectory
// 		storeDir = configData.StoreDirectory
// 		fileTypes = configData.FileTypes
// 		tagName = configData.InsertTag
// 		if configData.Port != "auto" {
// 			port, err = strconv.Atoi(configData.Port)
// 			if err != nil {
// 				printErr("Error parsing port:"+ err.Error())
// 				return
// 			}
// 		} else {
// 			// scan available port
// 			for i  := 8080; ; i++ {
// 				listener, err := net.Listen("tcp", ":"+strconv.Itoa(i))
// 				if err != nil {
// 					continue
// 				}
// 				port = listener.Addr().(*net.TCPAddr).Port
// 				break
// 			}
// 		}
// 		reConfig = regexp.MustCompile(`<!--\s*`+ configData.InsertTag+`\s*(?P<name>\S+)\s*start-->`)
// 		fmt.Println("read config.json successful !")
// 	}
	
// 	//Initialize render configuration
// 	renderConfigByte := []byte{}
// 	renderConfig, err := os.Open("./renderConfig.json")
// 	if err != nil {	
// 		renderConfigByte = []byte("{}")
// 		//create renderConfig.json if it doesn't exist
// 		renderConfig, err = os.Create("./renderConfig.json")
// 		if err != nil {
// 			printErr("Error creating renderConfig.json" + err.Error())
// 		}
// 	} else {
// 		fmt.Println("read renderConfig.json successful !")
// 		renderConfigByte, err = ioutil.ReadAll(renderConfig)
// 		if err != nil {
// 			printErr("Error reading file:"+ err.Error())
// 			return
// 		}
// 		renderConfigString := regexp.MustCompile(`(?m)^\s*//.*$|(?m)^\s*/\*[\s\S]*?\*/`).ReplaceAllString(string(renderConfigByte), "")
// 		structure := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(renderConfigString), &structure)
// 		if err != nil {
// 			printErr("Error parsing JSON:"+ err.Error())
// 			return
// 		}
// 		// Serialize map to JSON string
// 		renderConfigByte, err = json.Marshal(structure)
// 		if err != nil {
// 			printErr("Error parsing JSON:"+ err.Error())
// 			return
// 		}
// 		// fmt.Println(json.MarshalIndent(renderConfigByte, "", "    "))
// 		fmt.Println(string(renderConfigByte))

// 	}
// 	// Create map to store dataMap
// 	dataMap := make(map[string][]string)
// 	// Create map to store directory for each content
// 	dirs := make(map[string]string)
// 	//scan specified type of files in the directory
// 	filepath.Walk(scanDir, func(path string, info os.FileInfo, err error) error {
// 		// fileTypes = fileTypes.([]interface{})
// 		for _, fileType := range fileTypes {	
// 			if filepath.Ext(path) == fileType {
// 				// Read file contents
// 				b, err := ioutil.ReadFile(path)
// 				if err != nil {
// 					return err
// 				}
// 				// Extract "incert tag name" value from file contents
// 				matches := reConfig.FindAllStringSubmatch(string(b) ,-1)
// 				// Store file "incert tag name" values in map
// 				fileName := strings.TrimSuffix(filepath.Base(path), fileType)
// 				nameValues := make([]string, len(matches))
				
// 				if len(matches) > 0 {
// 					for i , match := range matches {
// 						nameValues[i] = match[1]
// 						dirs[fileName+"?"+match[1]] = path
// 					}
// 					// Check if the fileName is already in the dataMap
// 					uniqueName := fileName
// 					for i := 1; ; i++ {
// 						if _, ok := dataMap[uniqueName]; !ok {
// 							// Unique fileName found, break the loop
// 							break
// 						}
// 						// Append a number to the fileName and check again
// 						uniqueName = fileName + strconv.Itoa(i)
// 					}
// 					dataMap[uniqueName] = nameValues
// 				}

// 			}
// 		}
// 		return nil
// 	})

// 	// Serialize map to JSON string
// 	listData, err := json.MarshalIndent(dataMap, "", "    ")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
	
// 	fmt.Println(string(listData) +"\n")
// 	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println()
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		fmt.Fprintf(w, string(listData))
// 		// fmt.Println("some one visit homepage and fetched edit list")
// 	})
// 	http.HandleFunc("/getdata", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println()
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
// 		type DataMap struct {
// 			Page  string
// 			Content string
// 		}
// 		var dataMap DataMap

// 		b := make([]byte, r.ContentLength)
// 		r.Body.Read(b)
// 		err := json.Unmarshal([]byte(string(b)), &dataMap)
// 		if err != nil {
// 			printErr(err.Error())
// 			return
// 		}

// 		fmt.Println("editing",dataMap.Content,"on",dataMap.Page)
// 		//read json data and create it if it doesn't exist
// 		dir := storeDir+"/"+ dataMap.Page +"/"+ dataMap.Content +".json" 
// 		err = os.MkdirAll(storeDir+"/"+ dataMap.Page , os.ModePerm)
// 		if err != nil {
// 			printErr("Error creating directory:"+ err.Error())
// 			return
// 		}
// 		//read or create json file
// 		jsonFile, err := os.Open(dir)
// 		if err != nil {
// 			fmt.Println("File doesn't exist, creating it...")
// 			// File doesn't exist, create it
// 			jsonFile, err = os.Create(dir)
// 			if err != nil {
// 				printErr("Error creating file:"+ err.Error())
// 				return
// 			}
// 			// Write the JSON string to the file
// 			jsonString := "{\"type\":\"doc\",\"content\":[{\"type\":\"paragraph\",\"content\":[{\"type\":\"text\",\"text\":\"Example Text\"}]}]}"
// 			err = ioutil.WriteFile(dir, []byte(jsonString), 0644)
// 			if err != nil {
// 				printErr("Error writing to file:"+ err.Error())
// 				return
// 			}	
// 			fmt.Println(dir+" created")
// 		}
// 		defer jsonFile.Close()
// 		byteValue, err := ioutil.ReadAll(jsonFile)
// 		if err != nil {
// 			printErr("Error reading file:"+ err.Error())
// 			return
// 		}
// 		// Parse the JSON data
// 		var datas map[string]interface{}
// 		err = json.Unmarshal(byteValue, &datas)
// 		if err != nil {
// 			printErr("Error parsing JSON:"+ err.Error())
// 			return
// 		}
// 		fmt.Fprintf(w, string(byteValue))
// 		printSuccess("read from file "+dir+" successful")

// 	})
// 	http.HandleFunc("/savedata", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println()
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		//if request is Preflight, return
// 		if r.Method == "OPTIONS" {
// 			return
// 		}
// 		type Data struct {
// 			Page  string
// 			Content string
// 			Contentjson json.RawMessage
// 			Contenthtml string
// 		}
// 		var data Data
// 		//read json data
// 		b := make([]byte, r.ContentLength)
// 		r.Body.Read(b)
// 		err := json.Unmarshal((b), &data)
// 		if err != nil {
// 			// printErr(string(b))
// 			printErr(err.Error())
// 			return
// 		}
// 		fmt.Printf("saving to file %s\n")

// 		//Store json data
// 		dir := storeDir+"/"+data.Page+"/"+data.Content+".json" 
// 		jsonFile, err := os.Open(dir)
// 		if err != nil {
// 			printErr("Error saving file:"+ err.Error())
// 			return
// 		}
// 		defer jsonFile.Close()
// 		fmt.Printf("saving to file %s\n", dir)

// 		err = ioutil.WriteFile(dir, []byte(string(data.Contentjson)), 0644)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		//incert html data to html file
// 		// Open the file using the path stored in the dirs map
// 		file, err := os.Open(dirs[data.Page+"?"+data.Content])
// 		if err != nil {
// 			printErr("Error opening file:"+ err.Error())
// 			return
// 		}
// 		defer file.Close()
// 		// Read the file contents
// 		b, err = ioutil.ReadAll(file)
// 		if err != nil {
// 			printErr("Error reading file:"+ err.Error())
// 			return
// 		}
// 		// Use regular expressions to find the start and end tags
// 		startTag := regexp.MustCompile(`<!--\s*` + tagName +`\s*` + data.Content + `\s*start-->`)
// 		endTag := regexp.MustCompile(`<!--\s*` + tagName +`\s*` + data.Content+ `\s*end-->`)
// 		// Replace the contents between the tags with data.Contenthtml
// 		startPattern  := startTag.FindIndex(b)
// 		tagBefore := string(b[0:startPattern[1]])
// 		endPattern := endTag.FindIndex(b)
// 		var tagAfter string
// 		if len(endPattern) == 0 {
// 			tagAfter = "<!-- " + tagName +" " + data.Content+ " end-->\n" + string(b[startPattern[1] + 1:])
// 		}else{
// 			tagAfter = string(b[endPattern[0]:])
// 		}
// 		newContents := []byte(tagBefore + data.Contenthtml + tagAfter)
// 		// Write the modified contents back to the file
// 		err = ioutil.WriteFile(dirs[data.Page+"?"+data.Content], newContents, 0644)
// 		if err != nil {
// 			printErr("Error writing to file:"+ err.Error())
// 			return
// 		}
// 		printSuccess("saved " + dir + " successfully")
// 		fmt.Fprintf(w, "success")
// 	})
// 	http.HandleFunc("/getRenderconfig", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println()
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		fmt.Fprintf(w, string(renderConfigByte))
// 	})
// 	//check if running in production mode
// 	stat , err := os.Stat("./index.html")
// 	if stat != nil {
// 		// fmt.Println("Running in production mode")
// 		http.Handle("/", http.FileServer(http.Dir("./")))
// 	} else {
// 		// fmt.Println("Running in development mode")
// 		http.Handle("/", http.FileServer(http.Dir("../dist")))
// 	}
// 	printSuccess("Server started on port" + strconv.Itoa(port))
// 	fmt.Println("Local:\t\t"+"\033[0;36m"+"http://127.0.0.1:"+strconv.Itoa(port)+"/"+ "\033[0m")
// 	//get local ip
// 	if ip := getOutboundIP(); ip != nil {
// 		fmt.Println("Network:\t"+"\033[0;36m"+"http://"+ip.String()+":"+strconv.Itoa(port)+"/"+ "\033[0m")
// 	}
	
// 	err = http.ListenAndServe(":"+ strconv.Itoa(port), nil)
// 	if err != nil {
// 		printErr("Error starting server:"+ err.Error())
// 		return
// 	} 
// }

// getOutboundIP gets the preferred outbound ip of this machine
// func getOutboundIP () net.IP {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		printErr("Error getting local IP")
// 		return nil
// 	}
// 	defer conn.Close()
// 	localAddr := conn.LocalAddr().(*net.UDPAddr)
// 	return localAddr.IP
// } 

// func printErr(err string) {
// 	fmt.Println("\033[0;31m" , err , "\033[0m")
// }
// func printSuccess(msg string) {
// 	fmt.Println("\033[0;32m" , msg , "\033[0m")
// }

// `<!-- WikiBreeze {{name}} start-->`
