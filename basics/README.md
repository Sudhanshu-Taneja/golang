# golang

Go is a high-level general purpose programming language that is statically typed and compiled. It is known for the simplicity of its syntax and the efficiency of development that it enables by the inclusion of a large standard library supplying many needs for common projects.

It is a grabage collected language and thus, it has automated memory management as part of runtime.

In terms of execution speed, go is faster than JS, python, ruby & php. In terms of compilation speed, it's faster than rust, C, C++, java & C#.

# Types of variables - 
1. String
2. Boolean
3. Integer - Whole numbers
4. Unsigned integer - Only positive numbers
5. Float
6. Complex

# Functions

A function is a block of statements that can be used repeatedly in a program. It will not execute automatically when a page loads. It will be executed by a call to the function.

Create a Function - To create (often referred to as declare) a function, do the following:

1. Use the func keyword.
2. Specify a name for the function, followed by parentheses () where you can define the variables that the funciton takes as input and their types.
3. Finally, add code that defines what the function should do, inside curly braces {}.

Call a Function
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

# Defer 

The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers and the like.

# Closure

A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

# Structs

Structs are used to represent structured data. It's often convenient to group different types of variables together. 

# Interface

Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures. In our example, we've constructed a shape interface with 2 functions, area & perimeter. Now, any object that supports these 2 functions is an interface of "shape".

Type assertions - In GO, you can cast an interface to its underlying type using a type assertion.

Difference between an interface & classes - 

1. Interface don't have constructors or destructors that require that data is created or destroyed.
2. Interfaces aren't hierarchial by nature.
3. Interfaces define function signatures, but no underlying behaviour.


# Error handling

Error in GO is just an error interface which just has a function that returns a string.

# Continue keyword 

It stops the current iteration of a loop and continues to the next iteration.

# Break keyword

It stops the current iteration of a loop and exits the loop.

# Arrays

Arrays are stored in square brackets same as python. But, unlike python, the size of arrays in go is fixed.

# Slices

A slice a dynamically sized, flexible view into an array. We can use append function to dynamically add elements to a slice.

# Variadic functions

A variadic function is a function that accepts a variable number of arguments. This is achieved using an ellipsis (...) before the type of the parameter in the function signature.

# Spread operator

It allows us to pass a slice into a variadic function. It consists of 3 dots following the slice in the function call.

# Maps

It is very similar to dictionaries in python. Maps in GO are used to store data in key-value pairs. Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original.

# Higher order function

A function is called higher order function if it fulfills one of the below conditions - 

1. Passing function as an argument to another function - If a function is passed as an argument to another function, then such types of functions are known as higher-order function. It is also known as callback or first-class function.

2. Returning function from another function - If a function returns another function, then such type of function is known as higher order function. It is also known as first class function.

# Defer

It allows a function to be executed automatically just before its enclosing function returns. The deferred functions are typically used as cleaning functions performing tasks like closing database connections, file handlers etc.

# Closures

It is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

# Pointers

It is a variable that is used to store the memory address of another variable. The "&" is used to reference a variable to a pointer while "*" is used to de-reference a pointer and fetch it's value.

If a pointer points to nothing, then dereferencing it will cause a run-time error.

# Concurrency

It is the ability to perform multiple tasks at the same time. Typically, the code is executed one line at a time, one after the other. This is called sequential execution or synchronous execution.

If the computer we're running our code on has multiple cores, we can even execute multiple tasks at exactly the same time. Go was designed to be concurrent & it excels at performing many tasks simultaneously safely using the "go" keyword when calling a function. But, we can't capture any return statements while using this function.

# Channels

Channels are a typed, thread-safe queue. It allows different go-routines to communicate with each other. You can send and receive data from a channel using "<-" operator. These operations are blocking operations, which means if the channel is trying to send an output but there's no go routine to fetch that data, then the execution will stop until the channel finds a routine to send the output.

Blocking and deadlock - A deadlock is when a group of go-routines are all blocking so none of them can continue.

Channels in go can be explicitly closed by a sender indicating that we're done using the channel.

# Buffered channels 

By default, channels will only accept sends if there's a corresponding receive which are ready to receive the sent value. Buffered channels allows to accept a limited number of values without a corresponding reciever for those values. Buffered channel is blocked only when the buffer is full. Similarly, receiving from buffered channel is blocked only when the buffer will be empty.

# Select

It is used to listen to multiple channels at the same time. It is similar to switch statement but for channels.

# Mutexes

It allows us to lock access to data. This ensures that we can control which go-routine can access certain data at which time.

The best use-case of mutex is maps, as they are not safe for concurrent use. If we've multiple go-routines accessing the same map, and at least one of them is writing to the map, we should lock the maps with a mutex.

# R/W Mutex

Maps are safe for concurrent read access, just not concurrent read/write or write/write access. A read/write mutex allows all the readers to access the map at the same time, but a write will still lock out all other readers and writers.

# Generics

Generics allows the users to use variables to refer to specific types. It allows to write abstract functions that drastically reduce code duplication. It can be used in libraries and packages and thus, the code can be directly imported to be used in many appliactions.

# Constraints

A constraint means a type contraint, it is used to contrain some type parameters. We could view contraints as types of types. It allows us to use generics but for a subset of types. 