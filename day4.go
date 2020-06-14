package main
import "fmt"

// create a struct of type person with age as characteristic
type person struct {
	age int
}

// create a function with receiver for the struct created
// the function takes a variable of type int and returns type person
func (p person) NewPerson(initialAge int) person {

    if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	}
	p.age = initialAge
	return p
}

// create a function that will do the calculation
func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

// function to count years
func (p person) yearPasses() person {
	//Increment the age of the person in here
    p.age++
	return p
}

func main() {

	var T, age int
	// use fmt package to scan the first input
    fmt.Scan(&T)
	// check the number of inputs entered
    for i := 0; i < T; i++ {
		// take first age input
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p = p.yearPasses()
        }
        p.amIOld()
        fmt.Println()
    }
}	