package main

func main() {
	names := [10]string{"Amir", "Mori", "Aria", "Amin", "Ali", "Arash", "Shahryar", "Hadi", "Sorush", "Kia"}
	searchKey := "Kia"

	for i, index := range names {
		if searchKey == index {
			println("found in index : ", i)
		}
	}
}
