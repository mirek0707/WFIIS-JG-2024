package main

import (
	"encoding/xml"
)

type Person struct {
	XMLName    xml.Name `xml:"person"`
	Id         int      `xml:"id"`
	FirstName  string   `xml:"firstName"`
	LastName   string   `xml:"lastName"`
	Age        int      `xml:"age"`
	Birth      Data     `xml:"birth"`
	Death      Data     `xml:"death"`
	Pesel      string   `xml:"pesel"`
	CreditCard string   `xml:"creditcard"`
	Gender     string   `xml:"gender"`
}

type Data struct {
	D, M, Y int
}

type People struct {
	XMLName xml.Name  `xml:"persons"`
	People  []*Person `xml:"person"`
}

func (e *People) ReadXml(b []byte) error {
	err := xml.Unmarshal(b, e)
	return err
}
