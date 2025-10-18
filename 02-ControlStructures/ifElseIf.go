package main

import "fmt"

func main() {
	const minUserSalary float64 = 5_600_000
	var UserSalary float64
	var taxPercent float64 = 0

	fmt.Printf("Enter your Salary :")
	fmt.Scanf("%f",&UserSalary)

	if UserSalary <= minUserSalary {
		taxPercent = 0
	} else if UserSalary > minUserSalary && UserSalary <= minUserSalary*2 {
		taxPercent = 0.05
	} else if UserSalary > minUserSalary*2 && UserSalary <= minUserSalary*3 {
		taxPercent = 0.07
	} else if UserSalary > minUserSalary*3 && UserSalary <= minUserSalary*4 {
		taxPercent = 0.10
	} else {
		taxPercent = 0.15
	}

	fmt.Printf("Your tax percent is: %.2f\n", taxPercent)
	fmt.Printf("Your tax is: %.2f\n", taxPercent*UserSalary)
	
	fmt.Printf("Your UserSalary is: %.2f\n", UserSalary-taxPercent*UserSalary)
}