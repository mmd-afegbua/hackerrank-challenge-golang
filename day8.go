package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"io"
)

func main() {
	lineInput := bufio.NewReaderSize(os.Stdin, 1024*1024)

	lenInput, err := strconv.Atoi(lineReader(lineInput))
	checkError(err)

	phoneBook := make(map[string]int)

	for i := 0; i < lenInput; i++ {
		tempInput := strings.Split(lineReader(lineInput), " ")
		name := tempInput[0]
		phoneNum, _ := strconv.Atoi(tempInput[1])
		phoneBook[name] = phoneNum
	}

	for i := 0; i < lenInput; i++ {
		queryInput := lineReader(lineInput)
		if queryInput == "" {
			break
		}

		if val, ok := phoneBook[queryInput]; ok {
			fmt.Printf("%s=%d\n", queryInput, val)

		} else {
			fmt.Println("Not found")
		}
	}
}


func lineReader(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return " "
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
