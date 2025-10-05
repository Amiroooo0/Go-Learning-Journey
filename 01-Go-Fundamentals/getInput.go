package main

import (
	"fmt"
	"strings"
)

func main() {
	name := getNameCLI()
	vowels := vowelCount(name)
	fmt.Printf("Your name have %d vowels", vowels)
}

func getNameCLI() (name string) {
	fmt.Printf("Enter your name : ")
	fmt.Scanf("%s", &name)

	for _, item := range name {
		if !(((item <= 'z') && (item >= 'a')) || ((item <= 'Z') && (item >= 'A'))) {
			return ""
		}
	}

	return name
}

func vowelCount(name string) (result int) {
	result = 0
	for _, v := range strings.ToLower(name) {
		switch v {
		case 'a':
			result++
		case 'o':
			result++
		case 'u':
			result++
		case 'y':
			result++
		case 'i':
			result++
		case 'e':
			result++
		}
	}
	return result
}
