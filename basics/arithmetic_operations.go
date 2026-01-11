package main

func operations() {
	var a, b int = 13, 3
	var result int

	result = a + b // Addition
	println("Addition:", result)
	result = a - b // Subtraction
	println("Subtraction:", result)
	result = a * b // Multiplication
	println("Multiplication:", result)
	var result2 = float64(a) / float64(b) // Division
	println("Division float:", result2)
	println("Division int:", a/b)

	result = a % b // Modulus
	println("Modulus:", result)

	result++ // Increment
	println("Increment:", result)
	result-- // Decrement
	println("Decrement:", result)

}
