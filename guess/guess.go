package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Please think of a number between ", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	for {
		guess := (low + high) / 2
		fmt.Println("I guess the nubmer is", guess)
		fmt.Println("is that :")
		fmt.Println("(a) too high ?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct ")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("yaay I won")
			break
		} else {
			fmt.Println("Invalid response, try again.")
		}
	}
}