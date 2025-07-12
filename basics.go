package main // Package decoration which builds the file into executable go program.

import (
	"fmt" // Importing fmt package from standard library in order to use it.
	"math"
)

func calc(x, y int) (int, int, int) {
	// The type of variables can declared together in synchronized manner.
	// This function takes two integers x and y, adds them, and returns the result.
	return (x + y), (x - y), (x * y) // Return the sum and difference of x and y.
}

type car struct {
	brand string // Field brand of type string.
	model string // Field model of type string.
	year  int    // Field year of type int.
	Wheel        // Field of type Wheel.
}

type Wheel struct {
	Radius   float64 // Field size of type int.
	Material string  // Field type of type string.
}

// Methods in structs
func (w Wheel) getArea() int {
	// Method to get the area of the wheel.
	return int(math.Pi * w.Radius * w.Radius) // Return the radius of the wheel.
}

// Mock db struct and methods
type db struct{}

func (d *db) Open(name string) (*db, error) {
	// Simulate opening a database connection
	return &db{}, nil
}

func (d *db) FetchUser() (string, error) {
	// Simulate fetching a username
	return "mock_username", nil
}

// Update GetUsername function to use the mock db
func GetUsername(dstName, srcName string) (username string, err error) {
	// Create a new db instance
	conn, _ := (&db{}).Open(srcName)

	// Close the connection *anywhere* the GetUsername function returns
	defer func() {
		// Simulate closing the connection
		fmt.Println("Connection closed")
	}()

	username, err = conn.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}

	// The defer statement is auto-executed if we return here
	return username, nil
}

// Closure function
func findSum() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {

	// Print a string to the console.
	add, sub, _ := calc(10, 7)             // Call the add function with arguments 10 and 7, and assign the result to add and sub but ignore the multipication output.
	fmt.Println("The sum is:", add)        // Print the result of the addition to the console.
	fmt.Println("The difference is:", sub) // Print the result of the subtraction to the console.

	mycar := car{ // Create a new car instance with specified fields.
		brand: "Toyota",  // Assign the brand field to "Toyota".
		model: "Corolla", // Assign the model field to "Corolla".
		year:  2020,      // Assign the year field to 2020.
		Wheel: Wheel{ // Assign the Frontwheel field to a new Wheel instance.
			Radius:   15,       // Assign the Radius field of Frontwheel to 15.
			Material: "Rubber", // Assign the Material field of Frontwheel to "Rubber".
		},
	}
	fmt.Println("Car details:")                     // Print a header for car details.
	fmt.Println("Brand:", mycar.brand)              // Print the brand of the car.
	fmt.Println("Frontwheel Radius:", mycar.Radius) // Print the radius of the car.

	fmt.Println("Area of wheel: ", mycar.getArea()) // Print the area of the wheel using the getArea method.

	sum := findSum()
	fmt.Println(sum(5))  // Call the closure function with argument 5 and print the result.
	fmt.Println(sum(10)) // Call the closure function with argument 10 and print the result.

}
