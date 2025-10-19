package main

import "fmt"

type p struct {
	Name   string
	Family string
	Age    int
}

func main() {

	persons := make(map[string]p)
	var personSlice []string // for keeping order (map does not have order)
	persons["1234567890"] = p{Name: "Ali", Family: "Hoessini", Age: 30}
	personSlice = append(personSlice, "1234567890")
	persons["4832425354"] = p{Name: "Reza", Family: "Rezaei", Age: 31}
	personSlice = append(personSlice, "4832425354")
	persons["8579456213"] = p{Name: "Milad", Family: "Mohammadi", Age: 31}
	personSlice = append(personSlice, "8579456213")
	persons["9285365246"] = p{Name: "Peyman", Family: "Rezvani", Age: 31}
	personSlice = append(personSlice, "9285365246")
	persons["6321459957"] = p{Name: "Amir", Family: "Hamidi", Age: 31}
	personSlice = append(personSlice, "6321459957")
	persons["4563214569"] = p{Name: "Amin", Family: "Golmahalle", Age: 31}
	personSlice = append(personSlice, "4563214569")

	// Length map
	fmt.Printf("%d\n", len(persons))

	// Integration map
	for _, person := range personSlice {
		fmt.Printf("%v\n", persons[person])
	}

	// Copy map
	persons2 := make(map[string]p)

	for key, value := range persons {
		persons2[key] = value
	}
}
