package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"sort"
	"syscall"

	"io"
	"log"
	"os"

	"golang.org/x/term"
)

func ReadFile(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	data, err := io.ReadAll(file) // data: []byte
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func CreateFile(users *Users) {
	user := User{Login: "test", Password: "test", Role: 0}
	users.Users = append(users.Users, &user)
	b := users.SaveToXml()
	err := os.WriteFile("users.xml", b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteFile(users *Users) {
	b := users.SaveToXml()
	err := os.WriteFile("users.xml", b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var users Users
	if _, err := os.Stat("users.xml"); errors.Is(err, os.ErrNotExist) {
		CreateFile(&users)
	} else {
		usersData := ReadFile("users.xml")
		err := users.ReadXml(usersData)
		if err != nil {
			fmt.Println(err)
		}
	}

	fileName := "encrypted.xml"
	cipherText, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	key := "The giraffes enter the wardrobe."
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}
	// fmt.Println("text from file:\n", string(plainText))

	var people People
	err = people.ReadXml(plainText)
	if err != nil {
		fmt.Println(err)
	}

	//Login
	fmt.Println("Enter login: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	username := in.Text()
	sort.Slice(users.Users, func(i, j int) bool {
		return users.Users[i].Login <= users.Users[j].Login
	})
	idx := sort.Search(len(users.Users), func(i int) bool {
		return string(users.Users[i].Login) >= username
	})
	if idx < len(users.Users) && users.Users[idx].Login == username {
		fmt.Println("Found:", users.Users[idx].Login)
		fmt.Println("Enter password: ")
		password, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println(err)
		}

		if string(password) == users.Users[idx].Password {
			fmt.Println("User login succesfully")
			if users.Users[idx].Role == 1 || users.Users[idx].Role == 8 {
				fmt.Println("Add person:")

			}
		} else {
			fmt.Println("Wrong password")
		}

	} else {
		fmt.Println("User not found ")
	}

	// fmt.Println()
	// users.Print()
	WriteFile(&users)
}
