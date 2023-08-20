package xml

import "exchange/pkg/currency/xsd/tools"

func Run()  {
	EncodeFromFile()
	tools.DecodeFromXmlToJSON()
}