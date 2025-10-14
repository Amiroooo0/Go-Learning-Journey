package main

import "fmt"

func main() {
	var factorial uint64 = 1
	var number uint

	fmt.Printf("enter a number : ")
	fmt.Scanf("%d",&number)

	switch number {
		case 15 :
			factorial *= 15
			fallthrough
		case 14 :
			factorial *= 14
			fallthrough
		case 13 :
			factorial *= 13
			fallthrough
		case 12 :
			factorial *= 12
			fallthrough
		case 11 :
			factorial *= 11
			fallthrough
		case 10 :
			factorial *= 10
			fallthrough
		case 9 :
			factorial *= 9
			fallthrough
		case 8 :
			factorial *= 8
			fallthrough
		case 7 :
			factorial *= 7
			fallthrough
		case 6 :
			factorial *= 6
			fallthrough
		case 5 :
			factorial *= 5
			fallthrough
		case 4 :
			factorial *= 4
			fallthrough
		case 3 :
			factorial *= 3
			fallthrough
		case 2 :
			factorial *= 2
			fallthrough
		case 1 :
			factorial *= 1
			fallthrough
		case 0 :
			factorial *= 1
		default :
			fmt.Printf("number is too big")
	}
	fmt.Printf("factorial is : !%d = %d" , number , factorial)
}