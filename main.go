package main

// we can import like this
// import(
// 	"errors"
// 	"fmt"
// )
import (
	"errors"
	"fmt"
)

// for error handleing

// custom function
func greeting() {
	fmt.Println("Inside greeting function")
}

func add(a int, b int) int {
	return a + b
}

// func divide(a float64, b float64)-->If multiple parameters have the same type, you only need to write the type once — for the last parameter.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Welcome,", name)
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Ans:", ans)
	}

	// basics
	number := 8
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println("This is odd")
	}
	// Task : Switch statement for day names
	day := 3
	fmt.Print("Day ", day, ": ")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// functions
	greeting()

	result := add(5, 6)
	fmt.Println(result)
}

//	func main() {
//		const PI = 3.14
//		var name string = "Puspo"
//		// name:="Rahun"
//		age := 25 // short-hand declaration
//		fmt.Println("Name:", name)
//		fmt.Println("Age:", age)
//		fmt.Println("PI:", PI)
//	}
