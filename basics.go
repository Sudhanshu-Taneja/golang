package main // Package decoration which builds the file into executable go program.

import "fmt" // Importing fmt package from standard library in order to use it.

func calc(x, y int) (int, int, int) { // The type of variables can declared together in synchronized manner.
	// This function takes two integers x and y, adds them, and returns the result.
	return (x + y), (x - y), (x * y) // Return the sum and difference of x and y.
}

type car struct {
	brand      string // Field brand of type string.
	model      string // Field model of type string.
	year       int    // Field year of type int.
	Frontwheel Wheel  // Field Frontwheel of type Wheel.
	Backwheel  Wheel  // Field Backwheel of type Wheel.
}

type Wheel struct {
	Radius   int    // Field size of type int.
	Material string // Field type of type string.
}

func main() {
	num := 10 // Declare a variable num and assign it the value 10.
	// Print the value of num to the console.
	fmt.Println(num)
	// Print a string to the console.
	add, sub, _ := calc(10, 7)             // Call the add function with arguments 10 and 7, and assign the result to add and sub but ignore the multipication output.
	fmt.Println("The sum is:", add)        // Print the result of the addition to the console.
	fmt.Println("The difference is:", sub) // Print the result of the subtraction to the console.

	mycar := car{ // Create a new car instance with specified fields.
		brand: "Toyota",  // Assign the brand field to "Toyota".
		model: "Corolla", // Assign the model field to "Corolla".
		year:  2020,      // Assign the year field to 2020.
		Frontwheel: Wheel{ // Assign the Frontwheel field to a new Wheel instance.
			Radius:   15,       // Assign the Radius field of Frontwheel to 15.
			Material: "Rubber", // Assign the Material field of Frontwheel to "Rubber".
		},
		Backwheel: Wheel{ // Assign the Backwheel field to a new Wheel instance.
			Radius:   15,       // Assign the Radius field of Backwheel to 15.
			Material: "Rubber", // Assign the Material field of Backwheel to "Rubber".
		},
	}
	fmt.Println("Car details:")        // Print a header for car details.
	fmt.Println("Brand:", mycar.brand) // Print the brand of the car.
	fmt.Println("Model:", mycar.model) // Print the model of the car.
}
