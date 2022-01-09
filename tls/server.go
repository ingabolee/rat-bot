package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

func main() {

	//sample self signed certificate

	certificate := []byte(`-----BEGIN CERTIFICATE-----
MIIC3jCCAkegAwIBAgIUb19Y0xkL4mc9sIJyQ9ZRa9QsExkwDQYJKoZIhvcNAQEL
BQAwgYAxCzAJBgNVBAYTAmV1MQ8wDQYDVQQIDAZzeWRuZXkxDzANBgNVBAcMBnN5
ZG5leTEPMA0GA1UECgwGc3lkbmV5MRAwDgYDVQQLDAdzZWN0aW9uMQwwCgYDVQQD
DANqYW4xHjAcBgkqhkiG9w0BCQEWD2phbkBjb21wYW55LmNvbTAeFw0yMDA3MjEx
MjIwMjdaFw0yMTA3MjExMjIwMjdaMIGAMQswCQYDVQQGEwJldTEPMA0GA1UECAwG
c3lkbmV5MQ8wDQYDVQQHDAZzeWRuZXkxDzANBgNVBAoMBnN5ZG5leTEQMA4GA1UE
CwwHc2VjdGlvbjEMMAoGA1UEAwwDamFuMR4wHAYJKoZIhvcNAQkBFg9qYW5AY29t
cGFueS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAN+XhX7NMZbWgzbX
Ggtt99Teqs6pERTmEL3gLiwJW+p+oABV1/wLbWd68mViHowz03gsgw4SICMswNst
W7+k5qohKjExmk6N2XxRKRsVT38L+Kf9g8SiYFSRUO4djG70264/CaD6FOwLIiP8
x0WeqV3DSg6L462JrAxF6x+N8YspAgMBAAGjUzBRMB0GA1UdDgQWBBSAmSMKa6El
/g7gz2b3oUQ3W5Bl2zAfBgNVHSMEGDAWgBSAmSMKa6El/g7gz2b3oUQ3W5Bl2zAP
BgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4GBAHi4WWihZee8pY4jtf5j
S7NgXbA/Nnczo+eJF6A/U3hodJkorURcdu2orBxs9kcecl9uWRPJr6une9JRlRWh
9DXMdoq4+CuP5LN+SRMxIoB9GtvUjhRpC1dVynY7GW67hQytTrKJzM9SbyVj/WO2
sCBEycvzXmnjQ3mlozaxwucT
-----END CERTIFICATE-----`)

	key := []byte(`-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAN+XhX7NMZbWgzbX
Ggtt99Teqs6pERTmEL3gLiwJW+p+oABV1/wLbWd68mViHowz03gsgw4SICMswNst
W7+k5qohKjExmk6N2XxRKRsVT38L+Kf9g8SiYFSRUO4djG70264/CaD6FOwLIiP8
x0WeqV3DSg6L462JrAxF6x+N8YspAgMBAAECgYBkG/RybK8aSRtgz3hiy67eCYBS
nVH/mG7AhQJHRz13RZCf9c+Jkxg978dd60ugHIg/UzaucyEefqguuiVNLij0BsWq
k1nXyyl/7EFOO4d3HqIp4Ly45QmYXQ4XOayW2CAE9L8rLDGqnVfA5HHc3zR83v3c
m9Imba5ubzR/YVbHcQJBAP+bOMDBBZwcSWRA9YZGW6jhQ4gifuKbXpYQWMdotFLn
nuL3G5sLEJ7rrtnAJYed+Nb6rPXX9+daziPZpRq2uy0CQQDf761pLSa8YwXXAa//
uP+9AuK4B3U/L1e5U6cF2/md1i2gaw6PKedL6WtJBlQjoXQFVYAf0Jc/mh3Y5bhp
ht1tAkBdI+m9S1jI9wHDV2xgXnj+A//Atpk359fCpPhEyaGT6DTcjaDwkUqgLk+L
p1nFnknTxIqMFwONuWgOZjukjVuNAkEAxmXjfi5phjg5AU9WbbqqoPvAgAjjgLJi
Byis7o0Ary0FSX3v7TjT2jaYPZ9kxhiR4PPqSsWUat4RGYwVATFiVQJATEqleE2N
c4Yb4WX+LwdE26QqjFMrH8IJi4T7qrm5vud4G9ALFbx+NwSLysrdAEI5VZiv/ref
nqS/Hex9KPBMFA==
-----END PRIVATE KEY-----`)

	cert, err := tls.X509KeyPair(certificate, key)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "0.0.0.0:12001"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
		}
		go handleClient(conn)
	}
}

func handleClient(connection net.Conn) {
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	newFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
}
