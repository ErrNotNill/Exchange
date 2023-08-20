package xml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var file = FileOpen()
var url = "http://www.cbr.ru/scripts/XML_daily.asp?date_req=%v/%v/%v" //XX.XX.XXXX (Day, Month, Year)

func EditDateFromUrl(url string) {
	fmt.Sprintf(url, time.Now())
}

func EncodeFromFile() {
	decoder := xml.NewDecoder(file)
	valute := make([]Valute, 0)
	for {
		token, err := decoder.Token()
		if err != nil {
			log.Printf(err.Error())
		}
		if token == nil {
			break
		}
		switch tp := token.(type) {
		case xml.StartElement:
			if tp.Name.Local == "Valute" {
				var val Valute
				decoder.DecodeElement(&val, &tp)
				valute = append(valute, val)
			}
		}
	}
	for i := range valute {
		fmt.Println(valute[i])
	}
}

func EncodeFromUrl() {
	var netClient = http.Client{
		Timeout: time.Second * 10,
	}
	res, _ := netClient.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(body)

	exchange := Valute{}

	jsonErr := json.Unmarshal(body, &exchange)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(exchange)
}

func FileOpen() *os.File {
	file, err := os.Open("./pkg/currency/xml/currency.xml")
	if err != nil {
		log.Printf(err.Error())
	}
	return file
}
