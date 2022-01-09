package main

import (
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	TakeScreenShot()

}
func TakeScreenShot() {
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Fatal(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d_%s.png", i, bounds.Dx(), bounds.Dy(), RandomString())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

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
