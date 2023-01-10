package utils

import (
	"fmt"
)

type config struct {
	scanDir string `json:"scanDirectory"`
	storeDir string `json:"storeDirectory"`
	Port string `json:"port"`
	insertTag string `json:"incertTag"`
	fileTypes []string `json:"fileType"`
}
// TreeNode represents a content in the tree
type TreeNode struct {
	Subtree []*pages
}
type pages struct {
	name string
	content []*TreeNode
}

func fff(){
	fmt.Println("fff")
}
