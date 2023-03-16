package utils

var DefaultConfig = `{
	// Directory containing the page to be modified (e.g. "D:\\github\\web\\src\\pages")
	"scanDirectory": "..\\",

	// Directory to store the edited page (e.g. "D:\\github\\web\\src\\WikiData")
	"storeDirectory": ".\\WikiData",

	// Port to be used (e.g. "8080" or "auto")
	"port": "auto",

	// File type to be scan (e.g. [".html",....])
	"fileType":[".html",".vue"]
}`

var DefaultRenderConfig = `{
    "paragraph": {
        "HTMLAttributes": {
            // "class": {
            //     "default": "paragraph"
            // }
        },
        "tag": "p"
    },
    "table" : {
        "HTMLAttributes": {
            // "class": {
            //     "default": "table"
            // },
            // ":tablenote": {
            //     "default": "table"
            // }
        },
        "tag": "Tablegotool"
    },
    "otherConfigurations":{
        "headingLevels": [1,2,3,4],
        "citation": "true",
        "tablePro": "true",             //table with note
        "tableMustContainNote": "true"  //table must contain note
    }
}`

var DefaultRenderJson = `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Example Text"}]}]}`

var TestPageJson = `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"testPage"}]}]}`
