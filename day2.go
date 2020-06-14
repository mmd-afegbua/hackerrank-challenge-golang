package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
	"strings"
)

//  dataInput returns a Reader
var dataInput = bufio.NewReaderSize(os.Stdin, 1024 * 1024)

func main() {
	// first Reader input
	mealCost, _ := strconv.ParseFloat(inputLine(), 64)
	// second Reader input
	tipPercent, _ := strconv.ParseFloat(inputLine(), 64)
	// third Reader input
	taxPercent, _ := strconv.ParseFloat(inputLine(), 64)

	n := calc(mealCost, tipPercent, taxPercent)
	fmt.Println(n)
}

func calc(mealCost float64, tipPercent float64, taxPercent float64) float64 {
	tip := mealCost * (tipPercent) / 100
	tax := mealCost * (taxPercent / 100)
	totalCost := mealCost + tip + tax
	return totalCost
	fmt.Println(math.Round(totalCost))

}

// function to read string into a variable
func inputLine() string {
	// input returns a string from ReadString function
	// the '\n' signifies that Enter will end the input
	input, _ := dataInput.ReadString('\n')
	// TrimSuffix returns a given string without the provided suffix
	input = strings.TrimSuffix(input, "\n")
	return input
}