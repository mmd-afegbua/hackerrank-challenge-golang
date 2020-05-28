package main

// we will have to import onl the packages we nned
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// I like declaring variables, I could just use := instead.
	var read string
	// bufio is used to read data from 'Keyboard' using Stdin
	// the data is stores a bufio.Reader
	in := bufio.NewReader(os.Stdin)
	// Returns whatever is entered, up to the point the user pressed Enter 
	// Stdin returns two values; a success or error
	// we don't need the errors, so we might just use blank identifier (_)
	read, _ = in.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(read)
}
/*
Alternate solution without using blank identifier
	func main() {
		in := bufio.NewReader(os.Stdin)
		read, err := in.ReadString('\n')
		log.Fatal(err) //Don't forget to import "log"
		fmt.Println("Hellow, World.")
		fmt.Println(read)
	}
 */
 