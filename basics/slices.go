package main

import "fmt"

func main() {
	var numbers []int
	fmt.Println(numbers)

	var numbers1 = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	slice := numbers1[1:4]
	fmt.Println("Slice from index 1 to 3:", slice)
}
