package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	homeDir, err := os.UserHomeDir()
	fileDir := homeDir + "/Library/Containers/com.amazon.Kindle/Data/Library/Application Support/Kindle/Cache/KindleSyncMetadataCache.xml"
	rpath, err := filepath.Rel(cwd, fileDir)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	file, err := os.Open(rpath)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(data))
}
