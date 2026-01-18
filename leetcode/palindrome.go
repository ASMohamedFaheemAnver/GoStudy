package leetcode

import (
	"strconv"
)

func IsPalindromeV2(x int) bool {
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	num := x
	var reversed int
	for num > 0 {
		reversed = reversed*10 + num%10
		num = num / 10
	}
	return reversed == x
}

func isPalindrome(x int) bool {
	var reversedXString string = ""
	for _, v := range strconv.Itoa(x) {
		reversedXString = string(v) + reversedXString
	}
	reversedX, err := strconv.Atoi(reversedXString)
	if err != nil {
		println("Error:", err.Error())
		return false
	}
	if reversedX == x {
		return true
	}
	return false
}
