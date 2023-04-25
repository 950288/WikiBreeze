package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Config struct {
	ScanDir   string   `json:"scanDirectory"`
	StoreDir  string   `json:"storeDirectory"`
	Port      string   `json:"port"`
	FileTypes []string `json:"fileType"`
}

func ReadConfig() (Config, error) {
	var configDataString string
	var configData Config
	configDir := "../WikibreezeData/config/config.json"
	fmt.Println("reading config.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		//create config.json if it doesn't exist
		//make config directory if it doesn't exist
		err = os.MkdirAll("../WikibreezeData/config", 0755)
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
		jsonBytes, err := io.ReadAll(configFile)
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

	// decode json to struct
	err = json.Unmarshal([]byte(configDataString), &configData)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}
	fmt.Println("config:\t", configData)
	return configData, nil
}
func ReadRenderConfig() (string, error) {
	configRender := make(map[string]interface{})
	var configRenderString string
	configDir := "../WikibreezeData/config/renderConfig.json"
	fmt.Println("reading renderConfig.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		//create renderConfig.json if it doesn't exist
		//make config directory if it doesn't exist
		err = os.MkdirAll("../WikibreezeData/config", 0755)
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
		jsonBytes, err := io.ReadAll(configFile)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading renderConfig.json: %w", err))
		}
		configRenderString = string(jsonBytes)
		fmt.Println("read renderConfig.json sccessfully")
	}
	configFile.Close()
	// remove comments
	comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)
	configRenderString = comment.ReplaceAllString(configRenderString, "")

	// decode json to struct
	err = json.Unmarshal([]byte(configRenderString), &configRender)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}

	//convert configRender to json string
	RenderString, err := json.Marshal(configRender)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}
	fmt.Println("renderConfig:\t" + string(RenderString))
	return string(RenderString), nil
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
		ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			log.Fatal(fmt.Errorf("error listening to port: %w", err))
		}
		ln.Close()
	} else {
		// scan available port
		for i := 8080; ; i++ {
			ln, err := net.Listen("tcp", ":"+strconv.Itoa(i))
			if err != nil {
				continue
			}
			port = ln.Addr().(*net.TCPAddr).Port
			ln.Close()
			break
		}
	}
	return port
}
func ScanFiles(ScanDir string, FileTypes []string) (map[string]string, []byte, error) {
	fmt.Println("scanning files in " + ScanDir)
	// Create a editor dataMap
	dataMap := make(map[string][]string)
	// Create map to store directory for each content
	// using [fileName+"?"content] as key
	dirs := make(map[string]string)
	// Create a regular expression to extract content in each page
	reConfig := regexp.MustCompile(`<!--\s*WikiBreeze\s*(?P<name>\S+)\s*start-->`)
	// Scan specified type of files in the directory
	filepath.Walk(ScanDir, func(path string, info os.FileInfo, err error) error {
		for _, fileType := range FileTypes {
			if filepath.Ext(path) == fileType {
				// Read file contents
				file, err := os.Open(path)
				if err != nil {
					PrintErr("error opening file" + path + ":" + err.Error())
				}
				defer file.Close()

				// Read file contents
				b, err := io.ReadAll(file)
				if err != nil {
					return err
				}
				fileName := strings.TrimSuffix(filepath.Base(path), fileType)
				// Extract content name from file contents
				matches := reConfig.FindAllStringSubmatch(string(b), -1)
				contents := make([]string, len(matches))
				if len(matches) > 0 {
					for i, match := range matches {
						contents[i] = match[1]
						dirs[fileName+"?"+match[1]] = path
					}
					// Check if the fileName is already in the dataMap
					uniqueName := fileName
					for i := 1; ; i++ {
						if _, ok := dataMap[uniqueName]; !ok {
							// Unique fileName found, break the loop
							break
						}
						// Append a number to the fileName and check again
						uniqueName = fileName + strconv.Itoa(i)
					}
					dataMap[uniqueName] = contents
				}

			}
		}
		return nil
	})
	fmt.Println("scanned files successfully")
	if len(dataMap) == 0 {
		fmt.Println("no files found")
		//create testPage.html
		configFile, err := os.Create("../testPage.html")
		if err != nil {
			log.Fatal(fmt.Errorf("error creating testPage.html: %w", err))
		}
		//write default content to testPage.html
		_, err = configFile.Write([]byte(TestPage))
		if err != nil {
			log.Fatal(fmt.Errorf("error writing to testPage.html: %w", err))
		}
		fmt.Println("created testPage.html")
		contents := make([]string, 1)
		contents[0] = "testContent"
		dataMap["testPage"] = contents
		dirs["testPage?testContent"] = "../testPage.html"
		// copy testContent.json to WikibreezeData directory
		err = os.MkdirAll("../WikibreezeData/WikiData/testPage", 0755)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating ../WikibreezeData/WikiData/testPage directory: %w", err))
		}

		src := "./testContent.json"
		dst := "../WikibreezeData/WikiData/testPage/testContent.json"
		fin, err := os.Open(src)
		if err != nil {
			log.Fatal(fmt.Errorf("error open to testContent.json: %w", err))
		}
		defer fin.Close()

		fout, err := os.Create(dst)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating ../WikibreezeData/WikiData/testPage/testContent.json: %w", err))
		}
		defer fout.Close()

		_, err = io.Copy(fout, fin)

		if err != nil {
			log.Fatal(fmt.Errorf("copy testContent.json to WikibreezeData directory: %w", err))
		}

		fmt.Println("created testContent.json")

	}
	// Serialize map to JSON string
	dataMapByte, err := json.MarshalIndent(dataMap, "", "    ")
	if err != nil {
		return nil, nil, fmt.Errorf("error serializing dataMap: %w", err)
	}
	fmt.Println("dataMap:\t" + string(dataMapByte))
	fmt.Println("dirs:\t" + fmt.Sprint(dirs))
	return dirs, dataMapByte, nil
}
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		PrintErr("Error getting local IP")
		return nil
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
func PrintErr(err string) {
	// fmt.Println("\033[31m", err, "\033[0m")
	color.Red(err)
}
func PrintSuccess(msg string) {
	// fmt.Println("\033[32m", msg, "\033[0m")
	color.Green(msg)
}
func Cyanf(msg string) string {
	return color.New(color.FgCyan).SprintfFunc()(msg)
}
