package main

import "fmt"

func main() {
	var students [5]string
	fmt.Printf("Hello,\nI am simple GOlang app. I have been created to print some values from my array.")
	students[0] = "Alice"
	students[1] = "Lisa"
	students[2] = "Johnny"
	students[3] = "Owen"
	students[4] = "Sara"
	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Students: %v \n", students[4])
	fmt.Printf("Number of students: %v\n", len(students))
}
