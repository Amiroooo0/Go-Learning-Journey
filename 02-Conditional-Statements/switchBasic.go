package main

import "fmt"

func main() {
	var score float64
	var grade string

	fmt.Printf("enter your score :")
	fmt.Scanf("%f",&score)

	switch {
		case score < 10 && score >= 0 :
			grade = "D"
		case score < 13 && score >= 10 :
			grade = "C"
		case score < 15 && score >= 13 :
			grade = "B"
		case score < 17 && score >= 15 :
			grade = "A"
		case score <= 20 && score >= 17 :
			grade = "A+"
	}

	fmt.Printf("your grade is : %s",grade)
}