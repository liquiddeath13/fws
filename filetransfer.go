package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func retrieveFile(conn net.Conn, bufferSize int64, fName string, fSize int64, mode string) bool {
	file, flag := createFile(fName, fSize, mode)
	if !flag {
		return false
	}
	flag = readNetStreamToFile(file, conn, bufferSize, fSize, mode)
	file.Close()
	if !flag {
		//write decompress!!!
		flag = true
	}
	return flag
}

func readNetStreamToFile(file *os.File, conn net.Conn, bufferSize int64, fSize int64, mode string) bool {
	var receivedBytes int64
	reader := bufio.NewReader(conn)
	alreadySz := int64(-1)
	if mode == "append" {
		alreadySz = fileSize(file)
		if alreadySz > 0 {
			log.Println("Режим дополнения: Смещаем указатель записи на", byteCountSI(alreadySz))
			receivedBytes = alreadySz
			file.Seek(alreadySz, 0)
			for i := int64(0); i < alreadySz; i++ {
				reader.ReadByte()
			}
		}
	}
	err := readNetStreamSz(file, reader, receivedBytes, bufferSize, fSize)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return mode == "append" ||
		(extractErr(io.CopyN(file, reader, (fSize-receivedBytes))) == nil &&
			extractErr(reader.Read(make([]byte, (receivedBytes+bufferSize)-fSize))) == nil)
}
