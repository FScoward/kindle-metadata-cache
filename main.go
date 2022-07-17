package main

import (
	"encoding/csv"
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
	ASIN    string    `xml:"ASIN"`
	Title   Title     `xml:"title"`
	Authors []Authors `xml:"authors"`
}

type Title struct {
	Pronunciation string `xml:"pronunciation,attr"`
	Title         string `xml:",chardata"`
}

type Authors struct {
	Author Author `xml:"author"`
}
type Author struct {
	Pronunciation string `xml:"pronunciation,attr"`
	Author        string `xml:",chardata"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	homeDir, _ := os.UserHomeDir()
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

	response := Response{}
	xml.Unmarshal(data, &response)
	// for _, metadata := range response.AddUpdateList.MetaData {
	// 	fmt.Println(metadata)
	// }

	// --- csvへ変換
	rowCount := len(response.AddUpdateList.MetaData)
	csvRecords := make([][]string, rowCount)
	for i, metadata := range response.AddUpdateList.MetaData {
		csvRecords[i] = make([]string, 3)
		csvRecords[i][0] = metadata.ASIN
		csvRecords[i][1] = metadata.Title.Title
	}

	// --- 書き込み
	f, err := os.Create("test.csv")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	w := csv.NewWriter(f)
	w.WriteAll(csvRecords)
}
