package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func extractAddr(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}

func fileSize(file *os.File) int64 {
	fileStat, err := file.Stat()
	if err != nil {
		return -1
	}
	return fileStat.Size()
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

func extractErr(v interface{}, e error) error {
	return e
}

func extractVal(v interface{}, e error) interface{} {
	return v
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

var xFCounter = 1

func createFile(fName string, fSize int64, mode string) (*os.File, bool) {
	if fileExists(fName) && mode != "append" {
		//NEEDFIX
		split0 := strings.Split(fName, "\\")
		split := strings.Split(split0[len(split0)-1], ".")
		split[0] += " (" + fmt.Sprint(xFCounter) + ")"
		split0[len(split0)-1] = strings.Join(split, ".")
		xFCounter++
		log.Println("Такой файл уже существует, пытаемся создать файл с именем", split0[len(split0)-1])
		return createFile(strings.Join(split0, "\\"), fSize, mode)
	}
	file, err := os.OpenFile(fName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return file, true
}

func readString(conn net.Conn, length int64) string {
	str := make([]byte, length)
	conn.Read(str)
	return strings.Trim(string(str), ":")
}

func strToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
