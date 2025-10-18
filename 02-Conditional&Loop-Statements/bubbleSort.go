package main

import "fmt"

func main() {
	lst := []int{66, -6, -6, 2, 45623, 0, 21, -75, 50, 8}
	BubbleSort(lst)
	fmt.Printf("%v", lst)
}

func BubbleSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		if !IsSorted(list) {
			for j := 0; j < len(list)-1-i; j++ {
				if list[j] > list[j+1] {
					Swap(&list[j], &list[j+1])
				}
			}
		}
	}
}

func IsSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
