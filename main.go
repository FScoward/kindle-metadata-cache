package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Response struct {
	XMLName        xml.Name       `xml:"response"`
	SyncTime       string         `xml:"sync_time"`
	CacheMeataData CacheMeataData `xml:"cache_metadata"`
	AddUpdateList  AddUpdateList  `xml:"add_update_list"`
}
type CacheMeataData struct {
	Version string `xml:"version"`
}
type AddUpdateList struct {
	MetaData []MetaData `xml:"meta_data"`
}
type MetaData struct {
	ASIN  string `xml:"ASIN"`
	Title Title  `xml:"title"`
}
type Title struct {
	Pronunciation string `xml:"pronunciation,attr"`
	Title         string `xml:",chardata"`
}

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
	// fmt.Println(string(data))

	response := Response{}
	xml.Unmarshal(data, &response)
	fmt.Println(response)
}
