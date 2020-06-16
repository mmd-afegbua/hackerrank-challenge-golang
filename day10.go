package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"os"
	"bufio"
	"io"
)

func main() {
	inputLine := bufio.NewReader(os.Stdin)
	input := lineReader(inputLine)
	num, err := strconv.ParseInt(input, 10, 64)
	checkError(err)

	binNum := strconv.FormatInt(num, 2)
	binArray := strings.Split(binNum, "0")
	// max is float64 type
	var count float64

	for _, value := range binArray {
		splitNum, _ := strconv.ParseFloat(value, 64)
		count = math.Max(splitNum, count)
	}

	consecs := strconv.Itoa(int(count))
	fmt.Println(len(consecs))

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