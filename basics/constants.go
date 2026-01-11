package main

const PI = 3.14
const GRAVITY = 9.8

func constants() {
	println("Value of PI:", PI)
	println("Value of GRAVITY:", GRAVITY)
	// PI = 3.14159 // This will cause a compile-time error
	// GRAVITY = 9.81 // This will also cause a compile-time error

	const (
		monthDays = 30
		weekDays  = 7
	)

	println("Value of monthDays:", monthDays)
	println("Value of weekDays:", weekDays)
}
