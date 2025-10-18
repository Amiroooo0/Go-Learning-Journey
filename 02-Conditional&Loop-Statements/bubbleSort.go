package main

import "fmt"

func main() {
	lst := []int{96, 4321, -85, -85, 78, -6, 0}
	BubbleSort(lst)
	fmt.Printf("%v", lst)
}

func BubbleSort(list []int) {
	finished := false
	index := 0
	for !finished {
		for i, item := range list {
			if item > list[i+1] {
				temp := i
				i = list[i+1]
				list[i+1] = temp
				for j := i + 1; j < len(list)-1; j++ {
					temp := i
					i = list[i+1]
					list[i+1] = temp
				}
			}
			index = i
		}
		if index == len(list)-1 {
			finished = true
		} else {
			finished = false
		}
	}
}

/*
96 , 4321 , -85 , -85 , 78 , -6 , 0		1  step

*96 , 4321 , -85 , -85 , 78 , -6 , 0	2  step

96 ,-85 , -85 , 78 , -6 , 0 , *4321		3  step

-85 , -85 , 78 , -6 , 0 , *96 , 4321	4  step

*-85 , -85 , 78 , -6 , 0 , 96 , 4321	5  step

-85 , *-85 , 78 , -6 , 0 , 96 , 4321	6  step

-85 , -85 , -6 , 0 , *78 , 96 , 4321	7  step

*-85 , -85 , -6 , 0 , 78 , 96 , 4321	8  step

-85 , *-85 , -6 , 0 , 78 , 96 , 4321	9  step

-85 , -85 , *-6 , 0 , 78 , 96 , 4321	10 step

-85 , -85 , -6 , *0 , 78 , 96 , 4321	11 step

-85 , -85 , -6 , 0 , *78 , 96 , 4321	12 step

-85 , -85 , -6 , 0 , 78 , *96 , 4321	13 step

-85 , -85 , -6 , 0 , 78 , 96 , *4321	14 step
*/
