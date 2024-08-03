// run file using command: "go run main.go"

package main

import "fmt"

func main() {
	// declare n variable
	var n int

	// set the value of n
	n = 100

	// iterate through numbers from 1 to n
	for i := 1; i <= n; i++ {
		// check if the n number is divisible by both 3 and 5
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 { // check if the n number is divisible by 3
			fmt.Println("Fizz")
		} else if i%5 == 0 { // check if the n number is divisible by 5
			fmt.Println("Buzz")
		} else {
			// ff the n number is not divisible by 3 or 5, print the number itself
			fmt.Println(i)
		}
	}
}
