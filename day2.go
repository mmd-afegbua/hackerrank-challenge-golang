package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
	"strings"
)

var dataInput = bufio.NewReaderSize(os.Stdin, 1024 * 1024)

func main() {

	mealCost, _ := strconv.ParseFloat(inputLine(), 64)

	tipPercent, _ := strconv.ParseFloat(inputLine(), 64)

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

func inputLine() string {
	input, _ := dataInput.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}