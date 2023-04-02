package main

import (
	"WikiBreeze/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	config, err := utils.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	RenderConfigString, err := utils.ReadRenderConfig()
	if err != nil {
		log.Fatal(err)
	}

	port := utils.ScanPort(config.Port)

	// Get the list of file directories to be edit
	// using [fileName+"?"content] as key
	dirs, dataMapByte, err := utils.ScanFiles(config.ScanDir, config.FileTypes)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/list", utils.HandlerFetchContentList(dataMapByte))
	http.HandleFunc("/getdata", utils.HandlerGetContent(config.StoreDir))
	http.HandleFunc("/savedata", utils.HandlerSaveContent(config.StoreDir, dirs))
	http.HandleFunc("/getRenderconfig", utils.HandlergetRenderconfig(RenderConfigString))

	//check if running in production mode
	stat, _ := os.Stat("./index.html")
	if stat != nil {
		// fmt.Println("Running in production mode")
		http.Handle("/", http.FileServer(http.Dir("./")))
	} else {
		// fmt.Println("Running in development mode")
		http.Handle("/", http.FileServer(http.Dir("../dist")))
	}
	utils.PrintSuccess("Server started on port " + strconv.Itoa(port))
	fmt.Println("Local:\t\t", utils.Cyanf("http://127.0.0.1:"+strconv.Itoa(port)+"/"))

	//Get local ip
	if ip := utils.GetOutboundIP(); ip != nil {
		fmt.Println("Network:\t" + utils.Cyanf("http://"+ip.String()+":"+strconv.Itoa(port)+"/"))
	}

	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
