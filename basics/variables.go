package main

var PI = 3.14

const CONST_PI = 3.14

func main() {
	var age int = 30
	var unAssignedInt int
	var name string = "Alice"
	println(age)
	println(name)
	println(unAssignedInt)

	age = 31
	unAssignedInt = 100
	println(age)
	println(unAssignedInt)
	tooth := 68
	println(tooth)

	// Inline comment
	/*
		Multi-line comment
	*/

	// Can't access nameScope here
	// println(nameScope)

	println(PI)
	println(CONST_PI)
	PI = 3.14159
	println(PI)

	// Override global scope with block scope
	var PI = 3.2
	println(PI)
	// CONST_PI = 3.14159 // This will cause an error
}

func printName() {
	nameScope := "Bob"
	println(nameScope)
}
