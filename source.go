package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func processBasesList(conn net.Conn) {
	mightBeContentLength := readString(conn, 64)
	length, err := strToInt(mightBeContentLength)
	if err != nil {
		log.Println(err)
		return
	}
	body := readString(conn, length)
	for index, row := range strings.Split(string(body), ";") {
		split := strings.Split(row, "|")
		baseName := split[0]
		basePath := split[1]
		baseID := split[2]
		println(fmt.Sprint(index+1)+")", baseName, "-", basePath, "(", baseID, ")")
	}
	readString(conn, 64)
}

func processFilesList(conn net.Conn) {
	fSize, err := strToInt(readString(conn, 64))
	if err != nil {
		log.Println(err)
		return
	}
	remoteIP := extractAddr(conn)
	ensureDir(remoteIP)
	fName := remoteIP + "\\" + readString(conn, 64)
	mode := readString(conn, 64)
	log.Println("Стартует приём файла", fName, "размером", byteCountSI(fSize), "в режиме", mode)
	if retrieveFile(conn, 4096, fName, fSize, mode) {
		log.Println("Файл", fName, "успешно сохранён!")
	} else {
		log.Println("При приёме файла возникла ошибка!")
	}
}

func main() {
	//think about gzip
	ln, err := net.Listen("tcp", ":13370")
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := ln.Accept()
		go func(conn net.Conn) {
			println(extractAddr(conn) + ":")
			defer conn.Close()
			processFilesList(conn)
		}(conn)
	}
}
