package xml

import (
	"fmt"
	"net/http"
)

func GetXML() {
	request, _ := http.Get(url)
	body := request.Body
	fmt.Println(body)
}
