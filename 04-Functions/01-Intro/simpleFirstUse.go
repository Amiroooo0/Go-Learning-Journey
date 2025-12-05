package main

import "fmt"

func main() {
	str1 := "Hello World i2v\n"
	str2 := "Hello World"

	voidTOvoidPrint()
	println(voidToOutputPrint())
	inputToVoidPrint(str1)
	println(inputToOutputPrint(str2))
}

// without input and output
func voidTOvoidPrint() {
	fmt.Printf("Hello World v2v\n")
}

// without input and with output
func voidToOutputPrint() string {
	return "Hello World v2o"
}

// with input and without output
func inputToVoidPrint(str string) {
	fmt.Printf(str)
}

// with input and output
func inputToOutputPrint(str string) string {
	return str + " i2o"
}
