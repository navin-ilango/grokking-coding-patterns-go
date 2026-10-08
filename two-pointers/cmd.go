package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []string{
		"609",
		"88",
		"962",
		"101",
		"123",
	}

	for i, num := range nums {
		fmt.Printf("%d.\tnum: %s\n", i+1, num)
		fmt.Printf("\n\tIs strobogrammatic: %v\n", isStrobogrammatic(num))
		fmt.Println(strings.Repeat("-", 100))
	}

	s := "A man, a plan, a canal: Panama"
	result := isPalindrome(s)
	fmt.Printf("\n\tIs Palindrome: %v\n", result)
	fmt.Println(strings.Repeat("-", 100))

}
