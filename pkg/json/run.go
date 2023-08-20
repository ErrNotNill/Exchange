package json

import (
	"fmt"
	"io/ioutil"
)

func Run(){
	file, _ := ioutil.ReadFile("C:/projects/exchange/decodedjson.json")
	fmt.Println(string(file))

	//Instead of decode JSON we can just print JSON file
	//DecodeFileJSON() == fmt.Println(string(file))

	DecodeFileJSON()
}