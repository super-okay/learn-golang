package dataStructures

import (
	"fmt"
)

func Arrays() {
	// arrays can only be made up of same data types
	// fixed size
	grades := [3]int{99, 92, 88}    // array of integers with 3 values
	grades2 := [...]int{99, 92, 88} // initializing array with default values, no need to specify size, can use ...
	var students [3]string
	students[0] = "Lisa"
	students[2] = "Bob"
	students[1] = "Egg"
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	// arrays are considered values in Golang, instead of pointer to values
	a := [...]int{1, 2, 3}
	b := a
	b[0] = 100
	c := &a // c is pointing to a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	// slices are arrays without the length specified, have underlying array
	// pointers by default
	x := []int{1, 2, 3, 4, 5}
	y := x
	y[0] = 100
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)

	// array slicing works same way as in python
	slice1 := x[1:3]
	fmt.Printf("slice1: %v\n", slice1)

	// use make function to create arrays/slices
	// type, length, capacity of underlying array
	egg := make([]int, 3, 100) // length 3, capacity 100
	fmt.Printf("egg: %v\n", egg)
	// use append(source, value) to add item to slice
}

func Maps() {
	// key value pair
	// order not guaranteed, just like Python
	// can use make() to create map, or literal definition
	statePopulations := map[string]int{
		"California": 39250000,
		"Texas":      27860000,
		"New York":   19740000,
	}
	fmt.Println(statePopulations)
	fmt.Printf("%v\n", statePopulations["New York"])
	statePopulations["Georgia"] = 1031000
	fmt.Printf("%v\n", statePopulations["Georgia"])

	// pass by reference
	// slices, maps cannot be keys in maps
	// can delete keys
	// use comma ok syntax to check that a key is valid (value, ok)
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)
}

// similar to object, collection of disparate data types
// pass by value
// no inheritance, but can use composition via embedding
type Car struct {
	carType string
	seats   int
	color   string
}

func Structs() {
	car1 := Car{
		carType: "SUV",
		seats:   8,
		color:   "silver",
	}
	fmt.Println(car1)

	// can "index" into struct with index, but not scalable b/c indices change
	// order matters
}
