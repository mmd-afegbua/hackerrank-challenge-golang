package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"io"
)

func wordSort(word string) {
	var oddString string
	var evenString string
	for i, x := range word {
		selector := string(x)
		if i%2 == 0 {
			evenString += selector
		} else {
			oddString += selector
		}		
	}
	fmt.Printf("%s %s\n", evenString, oddString)
}

func main() {
	inputReader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	
	lenInput, err := strconv.ParseInt(lineReader(inputReader), 10, 64)
	if err != nil {
		panic(err)
	}
	for i := 0; i < int(lenInput); i++ {
		wordInput := bufio.NewScanner(os.Stdin)
		wordInput.Scan()
		word := wordInput.Text()
		wordSort(word)
	}

}

func lineReader(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return " "
	}
	return strings.TrimRight(string(str), "\r\n")
}