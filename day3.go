package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    // reader returns a buffered Reader
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    // create a variable that returns NTemp of int64 and error
    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    // check error using checkError function
    checkError(err)
    // convert Ntemp to int32 
	N := int32(NTemp)
	if N % 2 == 1 || N % 2 == 0 && N >= 6 && N <= 20 {
		fmt.Println("Weird")
	} else if N % 2 == 0 || N % 2 == 0 && N <= 2 && N >= 5 {
		fmt.Println("Not Weird")
	}
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }
    // TrimSuffix returns a given string without the provided suffix
    // '\r\n' is to ensure it runs on windows and unix
    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
