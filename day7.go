package main

import (
	"fmt"
	"bufio"
	"io"
	"strings"
	"strconv"
	"os"
)

func main() {
	lineInput := bufio.NewReaderSize(os.Stdin, 1024*1024)

	lenInput, err := strconv.ParseInt(lineReader(lineInput), 10, 64)
	if err != nil {
		panic(err)
	}

	var xInput []int

	tempInput := strings.Split(lineReader(lineInput), " ")
	
	for i := 0; i < int(lenInput); i++ {
		
		if i > int(lenInput) {
			fmt.Println("Ensure your input matches your input lenth")
		}
		xInputTemp, err := strconv.ParseInt(tempInput[i], 10, 64)
		if err != nil {
			panic(err)
		}

		n := int(xInputTemp)
		xInput = append(xInput, n)
	}	
	for i := len(xInput); i > 0; i-- {
		fmt.Printf("%d ", xInput[i-1])
	}
}

func lineReader(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return " "
	}
	return strings.TrimRight(string(str), "\r\n")
}
