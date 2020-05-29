package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv" // to convert the input from the keyboard that are always strings
)

func main() {
	// the 'var _' is added to occupy the strconv package by Hackerrank
	// this is because the solution might not use it, and the code won't run untill all
	// imported packages are used.
	var _ = strconv.Itoa //
	// declare variable i, d and s as instructed
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "Hackerrank"	
	
	// bufio reads incoming input and stores it in a variable
	scanner := bufio.NewScanner(os.Stdin)

	// thw incoming data will need to be stored as variables
	var (
		first uint64
		double float64
		words string
	)
	// we will be looking out for three inputs with Scan()
	// first scan
	scanner.Scan()
	// We know the first input, it is will enter as a text but we will
	// need to convert it to Uint64
	first, _ = strconv.ParseUint(scanner.Text(), 10, 64) // the 10, 64; couldn't find their why

	// second scan
	scanner.Scan()
	// again, it will be a Text() input, but we need it as a float64 type
	double, _ = strconv.ParseFloat(scanner.Text(), 640)

	// third scan
	scanner.Scan()
	// The input is originally Text(), no need for conversion
	words = scanner.Text()

	// it is what it is
	fmt.Printf("%v\n", i + first)
	fmt.Printf("%.1f\n", d + double)
	fmt.Printf("%s %s\n", s, words)
}