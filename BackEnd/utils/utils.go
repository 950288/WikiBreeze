package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Config struct {
	ScanDir   string   `json:"scanDirectory"`
	StoreDir  string   `json:"storeDirectory"`
	Port      string   `json:"port"`
	InsertTag string   `json:"incertTag"`
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
	var configData Config
	configDir := "./config/config.json"
	fmt.Println("reading config.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		//create config.json if it doesn't exist
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
	}
	configFile.Close()
	configFile, err = os.Open(configDir)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening config.json: %w", err))
	}
	defer configFile.Close()

	// read json file
	jsonBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return configData, err
	}

	jsonString := string(jsonBytes)

	// remove comments
	comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)
	jsonString = comment.ReplaceAllString(jsonString, "")

	// decode json to struct
	err = json.Unmarshal([]byte(jsonString), &configData)
	if err != nil {
		fmt.Println(jsonString)
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}

	return configData, nil
}
