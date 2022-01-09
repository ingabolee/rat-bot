package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/alistanis/goenc/aes/gcm"
)

func Encrypt(s, t string) {
	//sample key
	key := []byte("49236179763802224376317049986239")

	text, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
		Encrypt(s, t)
	}
	plaintext := []byte(text)

	ciphertext, err := gcm.Encrypt(key, plaintext, 12)
	if err != nil {
		log.Fatal(err)
		Encrypt(s, t)
	}
	enc := base64.StdEncoding.EncodeToString(ciphertext)

	f, err := os.Create(t)
	if err != nil {
		log.Fatal(err)
		Encrypt(s, t)
	}

	_, err = io.Copy(f, bytes.NewReader([]byte(enc)))
	if err != nil {
		log.Fatal(err)
		Encrypt(s, t)
	}
}

func main() {
	Encrypt("SShotpqC8hultaCPt", "filezzz.txt")
}
