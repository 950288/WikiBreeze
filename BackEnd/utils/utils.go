package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func ParseJson(filename string, configString string) (string, error) {
	jsonMap := make(map[string]interface{})
	// decode json to struct
	err := json.Unmarshal([]byte(configString), &jsonMap)
	if err != nil {
		return "", fmt.Errorf("error decoding %s: %w", filename, err)
	}

	//convert configString to json string
	encoded, err := json.MarshalIndent(jsonMap, "", "    ")
	if err != nil {
		return "", fmt.Errorf("error encoding %s: %w", filename, err)
	}
	return string(encoded), nil
}
func ReadDefaultConfig() string {
	DefaultConfigFile, err := os.Open("./src/config.json")
	if err != nil {
		log.Fatal(fmt.Errorf("error opening ./src/config.json : %w", err))
	}
	defer DefaultConfigFile.Close()
	Bytes, err := io.ReadAll(DefaultConfigFile)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading config.json: %w", err))
	}
	fmt.Println("read ./src/config.json sccessfully")

	// remove comments
	comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)

	ConfigString := comment.ReplaceAllString(string(Bytes), "")

	ConfigString, err = ParseJson("./src/config.json", ConfigString)
	if err != nil {
		log.Fatal(fmt.Errorf("error parse config.json: %w", err))
	}
	return ConfigString
}

func ReadDefaultEditorConfig() string {
	DefaultEditorConfigFile, err := os.Open("./src/editorConfig.json")
	if err != nil {
		log.Fatal(fmt.Errorf("error opening ./src/editorConfig.json : %w", err))
	}
	defer DefaultEditorConfigFile.Close()
	Bytes, err := io.ReadAll(DefaultEditorConfigFile)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading editorConfig.json: %w", err))
	}
	fmt.Println("read ./src/editorConfig.json sccessfully")

	// remove comments
	comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)

	EditorConfigString := comment.ReplaceAllString(string(Bytes), "")

	EditorConfigString, err = ParseJson("./src/editorConfig.json", EditorConfigString)
	if err != nil {
		log.Fatal(fmt.Errorf("error parse editorConfig.json: %w", err))
	}
	return EditorConfigString
}

func ReadConfig() map[string]interface{} {
	var configDataString string
	var configData map[string]interface{}
	configDir := "../WikibreezeData/config/config.json"
	fmt.Println("reading config.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		fmt.Println("config.json doesn't exist, creating it")
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
		configDataString = ReadDefaultConfig()
		_, err = configFile.Write([]byte(configDataString))
		if err != nil {
			log.Fatal(fmt.Errorf("error writing to config.json: %w", err))
		}
		fmt.Println("created config.json")
	} else {
		defer configFile.Close()
		// read json file
		jsonBytes, err := io.ReadAll(configFile)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading config.json: %w", err))
		}
		configDataString = string(jsonBytes)

		// remove comments
		comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)
		configDataString = comment.ReplaceAllString(configDataString, "")

		fmt.Println("read config.json sccessfully")
	}
	configFile.Close()

	// decode json to struct
	err = json.Unmarshal([]byte(configDataString), &configData)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing config.json: %w", err))
	}
	fmt.Println("config:\t", configData)
	return configData
}
func ReadEditorConfig() string {
	var editorConfigString string
	configDir := "../WikibreezeData/config/editorConfig.json"
	fmt.Println("reading editorConfig.json")
	configFile, err := os.Open(configDir)
	if err != nil {
		fmt.Println("editorConfig.json doesn't exist, creating it")
		//create editorConfig.json if it doesn't exist
		//make config directory if it doesn't exist
		err = os.MkdirAll("../WikibreezeData/config", 0755)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating config directory: %w", err))
		}
		//create editorConfig.json
		configFile, err = os.Create(configDir)
		if err != nil {
			log.Fatal(fmt.Errorf("error creating editorConfig.json: %w", err))
		}
		//write default config to editorConfig.json
		editorConfigString = ReadDefaultEditorConfig()
		_, err = configFile.Write([]byte(editorConfigString))
		if err != nil {
			log.Fatal(fmt.Errorf("error writing to editorConfig.json: %w", err))
		}
		fmt.Println("created editorConfig.json")
	} else {
		defer configFile.Close()
		// read json file
		jsonBytes, err := io.ReadAll(configFile)
		if err != nil {
			log.Fatal(fmt.Errorf("error reading editorConfig.json: %w", err))
		}
		editorConfigString = string(jsonBytes)

		// remove comments
		comment := regexp.MustCompile(`(?s)//.*?\n|/\*.*?\*/`)
		editorConfigString = comment.ReplaceAllString(editorConfigString, "")
		fmt.Println("read editorConfig.json sccessfully")

		editorConfigString, err = ParseJson("editorConfig.json", editorConfigString)
		if err != nil {
			log.Fatal(fmt.Errorf("error parsing editorConfig.json: %w", err))
		}
	}
	configFile.Close()

	fmt.Println("editorConfig:\n" + editorConfigString)
	return editorConfigString
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
		for i := 5001; i <= 65535; i++ {
			fmt.Println("scanning port " + strconv.Itoa(i))
			ln, err := net.Listen("tcp", ":"+strconv.Itoa(i))
			if err != nil {
				fmt.Println("port " + strconv.Itoa(i) + " is not available")
				continue
			}
			ln.Close()
			port = i
			break
		}
		// ln, err := net.Listen("tcp", ":0") // listen on a random port
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// port = ln.Addr().(*net.TCPAddr).Port
		// ln.Close()

	}
	return port
}
func ScanFiles(ScanDir string, FileTypes []interface{}) (map[string]string, []byte) {
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
				fileName := strings.TrimSuffix(filepath.Base(path), fileType.(string))
				// Extract content name from file contents
				matches := reConfig.FindAllStringSubmatch(string(b), -1)
				contents := make([]string, len(matches))
				if len(matches) > 0 {
					// Check if the fileName is already in the dataMap
					uniqueName := fileName
					for i := 1; ; i++ {
						if _, ok := dataMap[uniqueName]; !ok {
							// Unique fileName found, break the loop
							break
						}
						// Append a number to the fileName and check again
						uniqueName = fileName + "'" + strconv.Itoa(i)
					}

					for i, match := range matches {
						contents[i] = match[1]
						dirs[uniqueName+"?"+match[1]] = path
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
		contents[0] = "content"
		dataMap["testPage"] = contents
		dirs["testPage?content"] = "../testPage.html"
	}
	// Serialize map to JSON string
	dataMapByte, err := json.MarshalIndent(dataMap, "", "    ")
	if err != nil {
		log.Fatal("error serializing dataMap: %w", err)
		return nil, nil
	}
	fmt.Println("dataMap:\t" + string(dataMapByte))
	fmt.Println("dirs:\t" + fmt.Sprint(dirs))
	return dirs, dataMapByte
}

func CheckRedirect(req *http.Request, via []*http.Request) error {
	u, err := url.Parse(req.URL.String())
	if err != nil {
		return err
	}

	q := u.Query()
	q.Set("redirected", "true")
	u.RawQuery = q.Encode()

	req.URL = u

	fmt.Print("..")

	return nil
}
func GetCookie(account map[string]interface{}) (*cookiejar.Jar, string, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	username, ok_username := account["username"]
	password, ok_password := account["password"]
	if !ok_username || !ok_password {
		PrintErr("error reading username or password")
		log.Fatal()
	}
	usernameString, ok_username := username.(string)
	passwordString, ok_password := password.(string)
	if !ok_username || !ok_password {
		PrintErr("username or password is not string")
		log.Fatal()
	}

	client := &http.Client{
		// Transport: ,
		CheckRedirect: CheckRedirect, //disable redirect
		Jar:           jar,
	}

	req, err := http.NewRequest("POST", "https://old.igem.org/Login2", strings.NewReader("return_to=&username="+usernameString+"&password="+passwordString+"&Login=Login"))
	if err != nil {
		return nil, "", err
	}

	fmt.Print("trying to login..")
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	fmt.Println()
	teamsStatus, _ := ioutil.ReadAll(resp.Body)

	if !strings.Contains(string(teamsStatus), "successfully") {
		fmt.Println()
		return nil, "", fmt.Errorf("error logging in, please check your username and password")
	}
	PrintSuccess("login successfully")

	// request team id
	req, err = http.NewRequest("GET", "https://old.igem.org/aj/session_info?use_my_cookie=1", nil)
	if err != nil {
		fmt.Println()
		return nil, "", err
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println()
		return nil, "", err
	}
	defer resp.Body.Close()

	teamsStatus, _ = ioutil.ReadAll(resp.Body)

	teamsStatusMarshalIndent, _ := ParseJson("teamname", string(teamsStatus))

	jsonMap := make(map[string]interface{})

	err = json.Unmarshal(teamsStatus, &jsonMap)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding team id: %w", err)
	}
	if err != nil {
		return nil, "", fmt.Errorf("error parsing team id: %w", err)
	}
	teams, ok := jsonMap["teams"].([]interface{})
	if !ok {
		return nil, "", fmt.Errorf("error getting team id")
	}
	var maxYear time.Time
	var teamID string
	for _, team := range teams {
		teamData := team.(map[string]interface{})
		teamYear := teamData["team_year"].(string)
		teamRoleStatus := teamData["team_role_status"].(string)

		t, err := time.Parse("2006", teamYear)
		if err != nil {
			continue
		}
		if t.After(maxYear) && teamRoleStatus == "Accepted" {
			teamID = teamData["team_id"].(string)
			maxYear = t
		}
	}
	if teamID == "" {
		fmt.Println()
		return nil, "", fmt.Errorf("error getting team id form :" + teamsStatusMarshalIndent + "\nplease check your account")
	}
	PrintSuccess("got team id: " + teamID + " successfully")
	return jar, "https://shim-s3.igem.org/v1/teams/" + teamID + "/wiki", nil
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		PrintErr("Error getting local IP, LAN editing not available")
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
func PrintInfo(msg string) {
	color.Yellow(msg)
}
func Cyanf(msg string) string {
	return color.New(color.FgCyan).SprintfFunc()(msg)
}
