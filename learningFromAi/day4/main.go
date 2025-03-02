package main

import (
	"fmt"
	"os"
)

func main() {
	// var numbers [5]int
	// numbers[0] = 1
	// numbers[1] = 4
	// Declare and initialize an array of integers
	numbers := []int{1, 2, 3, 4, 5}
	// modern way
	// fmt.Println("Moder way loop:")
	// for i, v := range numbers {
	// 	fmt.Println("Value:", v, "Index:", i)
	// }
	for _, num := range numbers {
		fmt.Println(num)
	}
	// for i := 0; i < len(numbers); i++ {
	// 	numbers[i]++
	// 	fmt.Println(numbers[i])
	// }
	// fmt.Println(numbers)

	// slice are dynamic array we can grow or shrink
	num := []int{10, 20, 30}
	num = append(num, 3)
	fmt.Println(num) // element add from end

	// Maps (Key-Value Pairs)
	// Maps are like dictionaries in Python — they store key-value pairs.
	student := map[string]int{
		"Rahim": 20,
		"Karim": 30,
	}
	fmt.Println(student["Karim"])

	// range is used to loop through arrays, slices, or maps.
	numb := []int{1, 2, 3}
	for i, v := range numb {
		fmt.Println("Index:", i, "Value:", v)
	}

	// go can read and write from files
	content := []byte("Hi there. I'm on writting on the example.txt file")
	os.WriteFile("example.txt", content, 0o644)
	fmt.Println("File written successfully done.")
}
