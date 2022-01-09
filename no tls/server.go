package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

func main() {
	service := ":22000"
	listener, err := net.Listen("tcp4", service)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())

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
