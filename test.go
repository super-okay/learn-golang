// strong, statically typed
// excellent community
// simple, fast compile times, garbage collected
// built-in concurrency
// compiles to standalone binaries
// following https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

package main

import (
	"fmt"
	"strconv" // converting int to string
)

// variables declared but not used --> compile time error
// reduces tech debt

// can define variables at package level
var egg string = "eggroll"

// can define bunch of variables with var block
var (
	brand string  = "Arc'teryx"
	model string  = "Atom LT"
	color string  = "bitters"
	price float32 = 259.00
)

// uppercase variables at package scope --> exported from package, globally visible
// lowercase variables at package scope --> package scope
// block level
// var name lengths should reflect variable life
// keep acronyms uppercase, ex: "HTTP", "URL"

func convert() {
	var i int = 10
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var s string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s, s)
}

func variables() {
	var i int = 42
	var j float32 = 100

	// can only use := on new var names
	k := 99 // inferred to be integer
	// i := 99 is invalid

	fmt.Println(i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Println(k)
	fmt.Println(egg)
}

func booleans() {
	// uninitialized vars default to "zero value", which is false for bools, 0 for nums
	var n bool = true
	a := 1 == 1
	b := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	// int8, int16, int32, int64
	// uint8, uint16, uint32
	// float32, float64
	// complex64, complex128
}

func math() {
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // defaults to int b/c ops on two ints must become int
	fmt.Println(a % b) // mod only available on int types, not float
	// can't do ops on two different types
}

func strings() {
	// strings are aliases for bytes, can be converted to []byte
	// can be concatenated with +
	// UTF-8, immutable
	s := "hello world"
	b := []byte(s)
	fmt.Printf("%v, %T\n", s[2], s[2]) // 108, uint8 b/c letter is a byte
	fmt.Printf("%v, %T\n", b, b)
}

func runes() {
	// rune is int32 alias, single quotes, UTF-32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}

// enumerated constant
const (
	a = iota
	b = iota
	c = iota
)

func constants() {
	// lowercase --> no export
	// uppercase --> export
	// immutable, can be shadowed
	// must be calculable at compile time
	// special symbol iota allows related constants to be created easily, starts at 0
	const myConst int = 100
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
}

func arrays() {
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

func maps() {
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

func structs() {
	car1 := Car{
		carType: "SUV",
		seats:   8,
		color:   "silver",
	}
	fmt.Println(car1)

	// can "index" into struct with index, but not scalable b/c indices change
	// order matters
}

func main() {
	// variables()
	// convert()
	// booleans()
	// math()
	// strings()
	// runes()
	// constants()
	// arrays()
	// maps()
	structs()
}
