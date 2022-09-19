package logic

import (
	"fmt"
)

// must start with uppercase so func is public
func IfStatement() {
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it !")
	}
}
