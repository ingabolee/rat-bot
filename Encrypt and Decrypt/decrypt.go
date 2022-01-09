package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/alistanis/goenc/aes/gcm"
)

func Dec(n int, s string) {
	//sample key
	key := []byte("49236179763802224376317049986239")

	text, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("Problem opening file, please try again.")
	}

	dec, err := base64.RawStdEncoding.DecodeString(string(text))

	plaintext, err := gcm.Decrypt(key, dec, 12)
	if err != nil {
		log.Fatal(err)
	}

	if n == 1 {
		f, err := os.Create("log.txt") //file.tar.gzip   //log.txt
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(f, bytes.NewReader(plaintext))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("file decrypted successfuly.")

	} else if n == 2 {
		f, err := os.Create("file.tar.rar") //file.tar.gzip   //log.txt
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(f, bytes.NewReader(plaintext))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("file decrypted successfuly.")

	} else {
		fmt.Println("Input invalid!!")
	}

}

func main() {
	fmt.Println("D3CRYPTOR. \n")
	time.Sleep(1 * time.Second)
	var fileType int
	var name string
	fmt.Println("What is the type of file? \n1. Text file/.txt \n2. Compressed tar/rar/zip/gzip file")
	fmt.Scanln(&fileType)
	fmt.Println("Enter the name of the file: ")
	fmt.Scanln(&name)
	Dec(fileType, name)
}
