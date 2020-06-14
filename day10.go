package main

import (
	"fmt"
)


func lineReader(reader *bufio.Reader) string {
	str, _, err = reader.ReadLine()
	if err == io.EOF {
		return " "
	}
	return strings.TrimRight(string(str), '\r\n')
}

func checkError(err error) {
	if err =! nil {
		panic(err)
	}
}