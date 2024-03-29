package main

import "fmt"

func main() {
	whatWasSaid, whatAge := say()
	fmt.Println(whatWasSaid, whatAge)
	changeUsingPointer(&whatAge)
	fmt.Println("Actually, I am", whatAge)
}

func say() (string, int) {
	return "Hello, I am", 33
}

func changeUsingPointer(i *int) {
	fmt.Println("Location is at", i)
	newValue := 34
	*i = newValue
}
