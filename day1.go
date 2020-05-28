package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _ = strconv.Itoa //
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "Hackerrank"	
	
	scanner := bufio.NewScanner(os.Stdin)

	var (
		first uint64
		double float64
		words string
	)

	scanner.Scan()
	first, _ = strconv.ParseUint(scanner.Text(), 10, 64)

	scanner.Scan()
	double, _ = strconv.ParseFloat(scanner.Text(), 640)

	scanner.Scan()
	words = scanner.Text()

	fmt.Printf("%v\n", i + first)
	fmt.Printf("%.1f\n", d + double)
	fmt.Printf("%s %s\n", s, words)
}