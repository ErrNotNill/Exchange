package tools

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs" json:"xml_name"`
	Date    string   `xml:"Date,attr" json:"date"`
	Name    string   `xml:"name,attr" json:"name"`
	Valute  []struct {
		ID       string `xml:"ID,attr" json:"ID"`
		NumCode  string `xml:"NumCode" json:"numCode"`
		CharCode string `xml:"CharCode" json:"charCode"`
		Nominal  string `xml:"Nominal" json:"nominal"`
		Name     string `xml:"Name" json:"name"`
		Value    string `xml:"Value" json:"value"`
	} `xml:"Valute"`
}

func DecodeFromXmlToJSON() {
	fContent, _ := ioutil.ReadFile("./pkg/currency/xml/currency.xml")
	var valute ValCurs
	xml.Unmarshal(fContent, &valute)
	jsonData, _ := json.MarshalIndent(valute, "", " ")
	dir, _ := ioutil.TempDir("C:/projects/exchange/pkg/currency/xml", "temp")

	ioutil.WriteFile("decodedjson.json", jsonData, 0644)
	fmt.Println(string(jsonData))
	defer os.RemoveAll(dir)
}

func FileOpen() *os.File {
	file, err := os.Open("./pkg/currency/xml/currency.xml")
	if err != nil {
		log.Printf(err.Error())
	}
	return file
}

var file = FileOpen()
