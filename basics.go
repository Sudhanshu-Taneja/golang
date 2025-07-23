package main // Package decoration which builds the file into executable go program.

import (
	"errors"
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

// Interface
type shape interface {
	// Define an interface with a method to calculate area
	area() float64
	perimeter() float64
}

// Circle struct implementing the shape interface
type Circle struct {
	radius float64 // Field radius of type float64.
}

// Implementing the area method for Circle
func (c Circle) area() float64 {
	// Calculate the area of the circle
	return math.Pi * c.radius * c.radius // Return the area of the circle.
}

// Implementing the perimeter method for Circle
func (c Circle) perimeter() float64 {
	// Calculate the perimeter of the circle
	return 2 * math.Pi * c.radius // Return the perimeter of the circle.
}

// Implementing variadic functions
func getSum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func getAddition(a, b int) int {
	return a + b // Return the sum of a and b.
}

func getProduct(a, b int) int {
	return a * b // Return the product of a and b.
}

// Implementing higher-order functions
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c) // Apply the arithmetic function to a and b, then to the result and c.
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
			Radius:   15,        // Assign the Radius field of Frontwheel to 15.
			Material: "Rubber ", // Assign the Material field of Frontwheel to "Rubber".
		},
	}
	fmt.Println("Car details:")                             // Print a header for car details.
	fmt.Println("Brand:", mycar.brand)                      // Print the brand of the car.
	s := fmt.Sprintf("Frontwheel Radius: %v", mycar.Radius) // Print the radius of the car.
	fmt.Println(s)                                          // Print the formatted string containing the radius of the front wheel.

	fmt.Println("Area of wheel: ", mycar.getArea()) // Print the area of the wheel using the getArea method.

	sum := findSum()
	fmt.Println(sum(5))  // Call the closure function with argument 5 and print the result.
	fmt.Println(sum(10)) // Call the closure function with argument 10 and print the result.

	//c, ok := s.(Circle) // Type assertion to check if s is of type Circle.
	/*
		"c" is a new circle cast from "s" which is of type "shape".
		"ok" is a boolean indicating whether the assertion was successful. If ok is true, then c is a Circle.
	*/

	// Error handling
	newSum, err := findSum(), error(nil) // Initialize newSum with the closure function and err with nil.
	if err != nil {
		fmt.Println("Error occurred:", err) // If there is an error, print it.
		return
	}
	fmt.Println("New sum is:", newSum(15)) // Print the result of the closure function with argument 15.

	// Creating dynamic error as per the code
	if newSum(0) == 15 {
		var err error = errors.New("Your sum value has reached 15. Testing new error creation.")
		fmt.Println(err)
	}

	// Creating loops
	for i := 0; ; i++ { // Loop from 0 to 4.
		fmt.Println("Loop iteration:", i) // Print the current iteration.
		if i == 4 {                       // If the iteration reaches 4, break the loop.
			fmt.Println("Breaking the loop at iteration:", i) // Print a message indicating the loop is breaking.
			break                                             // Break the loop.
		}
	}
	j := 5
	for j > 0 {
		fmt.Println("Loop iteration:", j) // Print the current iteration.
		j--                               // Decrement j by 1.
	}

	// Testing modulo operator and logical AND operator
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fiz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Arrays
	arr := [5]int{1, 2, 3, 4, 5}   // Declare an array of integers with 5 elements.
	fmt.Println("Array elements:") // Print a header for array elements.
	for i, v := range arr {        // Iterate over the array using range.
		fmt.Printf("Element at index %d: %d\n", i, v) // Print the index and value of each element in the array.
	}

	// Slices
	// func make([]T, len, cap)
	slice := arr[1:4]              // Declare a slice of integers with 5 elements.
	fmt.Println("Slice elements:") // Print a header for slice elements.
	for i, v := range slice {      // Iterate over the slice using range.
		fmt.Printf("Element at index %d: %d\n", i, v) // Print the index and value of each element in the slice.
	}

	sum1 := getSum(1, 2, 3)
	fmt.Println("Sum of numbers: ", sum1)

	sum2 := getSum(4, 7, 9, 11, 15, 19)
	fmt.Println("Sum of numbers: ", sum2)

	// Implementing spread operator
	sum3 := getSum([]int{1, 2, 3, 4, 5}...)                     // Using spread operator to pass a slice as variadic arguments.
	fmt.Println("Sum of numbers using spread operator: ", sum3) // Print the result of the sum using the spread operator.

	ages := map[string]int{ // Declare a map with string keys and int values.
		"Alice": 30, // Key "Alice" with value 30.
		"Bob":   25, // Key "Bob" with value 25.
	}
	ages["Charlie"] = 35          // Add a new key-value pair to the map.
	fmt.Println("Ages map:")      // Print a header for the ages map.
	for name, age := range ages { // Iterate over the map using range.
		fmt.Printf("%s is %d years old\n", name, age) // Print the name and age of each person in the map.
	}

	delete(ages, "Bob")       // Delete the key "Bob" from the map.
	elem, ok := ages["Alice"] // Check if the key "Alice" exists in the map.
	if ok {
		fmt.Println("Alice's age is:", elem) // Print Alice's age if the key exists.
	} else {
		fmt.Println("Alice not found in the map") // Print a message if the key does not exist.
	}

	// Using the aggregate function with addition and multiplication
	sumResult := aggregate(1, 2, 3, getAddition)
	fmt.Println("Sum result:", sumResult) // Print the result of the aggregation.

	productResult := aggregate(1, 2, 3, getProduct)
	fmt.Println("Product result:", productResult) // Print the result of the aggregation.

}
