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