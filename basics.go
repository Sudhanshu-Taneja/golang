package main // Package decoration which builds the file into executable go program.

import "fmt" // Importing fmt package from standard library in order to use it.

func add(x, y int) int { // The type of variables can declared together in synchronized manner.
	// This function takes two integers x and y, adds them, and returns the result.
	return x + y
}

func main() {
	num := 10 // Declare a variable num and assign it the value 10.
	// Print the value of num to the console.
	fmt.Println(num)
	// Print a string to the console.
	res := add(5, 7)                // Call the add function with arguments 5 and 7, and assign the result to res.
	fmt.Println("The sum is:", res) // Print the result of the addition to the console.
}
