package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi, I am a calculator app")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter calculation (Example: 2 + 2): ")
		scanner.Scan()
		input := scanner.Text()

		var a, b int
		var operator string

		fmt.Sscanf(input, "%d %s %d", &a, &operator, &b)

		switch operator {
		case "+":
			fmt.Println("Result:", a+b)
		case "-":
			fmt.Println("Result:", a-b)
		case "*":
			fmt.Println("Result:", a*b)
		case "/":
			fmt.Println("Result:", a/b)
		default:
			fmt.Println("Invalid operation")
		}
	}
}
