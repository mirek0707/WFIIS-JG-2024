package main

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []*User  `xml:"user"`
}

func (e Users) Print() {
	fmt.Println("Users:")
	for i := range e.Users {
		e.Users[i].Print()
	}
	fmt.Println()
}

func (e *Users) ReadXml(b []byte) error {
	err := xml.Unmarshal(b, e)
	for i := range e.Users {
		b, err := hex.DecodeString(e.Users[i].Password)
		if err != nil {
			fmt.Println(err)
		}
		e.Users[i].Password = string(b)
	}
	return err
}

func (e *Users) SaveToXml() []byte {
	usersCopy := e
	usersCopy.HashPasswords()
	b, err := xml.Marshal(usersCopy)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func (e *Users) HashPasswords() {
	for i := range e.Users {
		e.Users[i].Password = hex.EncodeToString([]byte(e.Users[i].Password))

		// b, err := hex.DecodeString(e.Users[i].Password)
		// s := string(b)
		// fmt.Println(s)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

}

type User struct {
	XMLName  xml.Name `xml:"user"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
	Role     int      `xml:"role"`
	// 0 - none
	// 1 - add
	// 2 - edit
	// 4 - read
	// 8 - all
}

func (e User) Print() {
	fmt.Println("\t", e.Login, "|", e.Password, "|", e.Role, "|", e.XMLName)
}
