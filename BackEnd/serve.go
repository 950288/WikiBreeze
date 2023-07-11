package main

import (
	"WikiBreeze/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/fatih/color"
)

func main() {
	config := utils.ReadConfig()

	RenderConfigString := utils.ReadEditorConfig()

	port := utils.ScanPort(config["port"].(string))

	// Get the list of file directories to be edit
	// using [fileName+"?"+content] as key
	dirs, dataMapByte := utils.ScanFiles(config["scanDirectory"].(string), config["fileType"].([]interface{}))

	fmt.Println()
	utils.ReadUploadConfig_HandleUpload(config)
	http.HandleFunc("/list", utils.HandlerFetchContentList(dataMapByte))
	http.HandleFunc("/getdata", utils.HandlerGetContent(utils.StoreDir))
	http.HandleFunc("/savedata", utils.HandlerSaveContent(utils.StoreDir, dirs))
	http.HandleFunc("/getEditorConfig", utils.HandlergetRenderconfig(RenderConfigString))

	//check if running in production mode
	stat, _ := os.Stat("./index.html")
	if stat != nil {
		// fmt.Println("Running in production mode")
		http.Handle("/", http.FileServer(http.Dir("./")))
	} else {
		// fmt.Println("Running in development mode")
		http.Handle("/", http.FileServer(http.Dir("../Wikibreeze")))
	}
	fmt.Println("")
	utils.PrintSuccess("WikiBreeze(v1.1.0-beta) started on port " + strconv.Itoa(port) + " successfully")
	fmt.Println("Local:\t\t", utils.Cyanf("http://localhost:"+strconv.Itoa(port)+"/"))

	//Get local ip
	if ip := utils.GetOutboundIP(); ip != nil {
		fmt.Println("Network:\t", utils.Cyanf("http://"+ip.String()+":"+strconv.Itoa(port)+"/"))
	}

	color.Magenta("Press CTRL+C to quit")

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}

	// listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server1 := http.Server{
	// 	Addr:    "localhost:" + strconv.Itoa(port),
	// 	Handler: nil,
	// }
	// go server1.Serve(listener)

	// server0 := http.Server{
	// 	Addr:    ":" + strconv.Itoa(port),
	// 	Handler: nil,
	// }
	// server0.Serve(listener)
}
