package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	ScanDir   string   `json:"scanDirectory"`
	StoreDir  string   `json:"storeDirectory"`
	Port      string   `json:"port"`
	TagName   string   `json:"tagName"`
	FileTypes []string `json:"fileType"`
}

// TreeNode represents a content in the tree
type TreeNode struct {
	Subtree []*pages
}
type pages struct {
	name    string
	content []*TreeNode
}

func ReadConfig() (Config, error) {
	var configDataString string
	var configData Config
	configDir := "./config/config.json"
	fmt.Println("reading config.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		//create config.json if it doesn't exist
		//make config directory if it doesn't exist
		err = os.MkdirAll("./config", 0755)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating config directory: %w", err))
		}
		//create config.json
		configFile, err = os.Create(configDir)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating config.json: %w", err))
		}
		//write default config to config.json
		_, err = configFile.Write([]byte(DefaultConfig))
		if err != nil {
			log.Fatal(fmt.Errorf("error writing to config.json: %w", err))
		}
		fmt.Println("created config.json")
		configDataString = DefaultConfig
	} else {
		defer configFile.Close()

		// read json file
		jsonBytes, err := ioutil.ReadAll(configFile)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading config.json: %w", err))
		}
		configDataString = string(jsonBytes)
		fmt.Println("read config.json sccessfully")
	}
	configFile.Close()

	// remove comments
	comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)
	configDataString = comment.ReplaceAllString(configDataString, "")
	fmt.Println(configDataString)

	// decode json to struct
	err = json.Unmarshal([]byte(configDataString), &configData)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}
	return configData, nil
}
func ReadRenderConfig() (string, error) {
	configRender := make(map[string]interface{})
	var configRenderString string
	configDir := "./config/renderConfig.json"
	fmt.Println("reading renderConfig.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		//create renderConfig.json if it doesn't exist
		//make config directory if it doesn't exist
		err = os.MkdirAll("./config", 0755)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating config directory: %w", err))
		}
		//create renderConfig.json
		configFile, err = os.Create(configDir)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating renderConfig.json: %w", err))
		}
		//write default config to renderConfig.json
		_, err = configFile.Write([]byte(DefaultRenderConfig))
		if err != nil {
			log.Fatal(fmt.Errorf("error writing to renderConfig.json: %w", err))
		}
		fmt.Println("created renderConfig.json")
		configRenderString = DefaultRenderConfig
	} else {
		defer configFile.Close()

		// read json file
		jsonBytes, err := ioutil.ReadAll(configFile)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading renderConfig.json: %w", err))
		}
		configRenderString = string(jsonBytes)
		fmt.Println("read config.json sccessfully")
	}
	configFile.Close()
	// remove comments
	comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)
	configRenderString = comment.ReplaceAllString(configRenderString, "")
	fmt.Println(configRenderString)

	// decode json to struct
	err = json.Unmarshal([]byte(configRenderString), &configRender)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}

	return configRenderString, nil
}
func ScanPort(Port string) int {
	var port int
	var err error
	if Port != "auto" {
		port, err = strconv.Atoi(Port)
		if err != nil {
			log.Fatal(fmt.Errorf("error parsing port: %w", err))
		}
		// check if port is available
		_, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			log.Fatal(fmt.Errorf("error listening to port: %w", err))
		}
	} else {
		// scan available port
		for i := 8080; ; i++ {
			listener, err := net.Listen("tcp", ":"+strconv.Itoa(i))
			if err != nil {
				continue
			}
			port = listener.Addr().(*net.TCPAddr).Port
			break
		}
	}
	return port
}
