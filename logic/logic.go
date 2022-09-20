package logic

import (
	"fmt"
)

// must start with uppercase so func is public
func IfStatement() {
	number := 50
	guess := 50
	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("You got it !")
	}
}

func SwitchStatement() {
	switch 3 {
	case 1, 3, 5:
		fmt.Println("one, three, five")
	case 2, 4, 6:
		fmt.Println("two, four, six")
	default:
		fmt.Println("none of the above")
	}

	switch i := 4 + 5; i {
	case 9:
		fmt.Println("yeah !")
	case 100:
		fmt.Println("no !")
	default:
		fmt.Println("default !")
	}
}

func ForLoop() {
	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// while does not exist in go, it is done with modified for loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	// break and continue work the same way as in Python
	// use labels to tell breaks which loop to break out of
	// looping over collections (arrays, slices, maps, strings): for k, v := range collection {}
}
