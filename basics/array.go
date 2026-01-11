package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)

	fmt.Println("Length of array:", len(numbers))

	originalArray := [3]string{"Go", "Python", "Java"}
	copiedArray := originalArray

	copiedArray[0] = "JavaScript"

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	for i := 0; i < len(originalArray); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, originalArray[i])
	}

	originalArray1 := [2]int{1, 2}
	copiedArray1 := &originalArray1

	fmt.Println("Original Array before pointer modification:", originalArray1)
	copiedArray1[0] = 10
	fmt.Println("Original Array after pointer modification:", originalArray1)
}
