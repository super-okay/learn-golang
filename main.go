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

	// package imports are relative to $GOPATH/src
	ds "../../Documents/learn-go/data-structures"
	logic "../../Documents/learn-go/logic"
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

func temp() {
	ds.Arrays()
	ds.Maps()
	ds.Structs()
	logic.IfStatement()
}

func main() {
	// variables()
	// convert()
	// booleans()
	// math()
	// strings()
	// runes()
	// constants()
	// ds.Arrays()
	// ds.Maps()
	// ds.Structs()
	logic.IfStatement()
}
