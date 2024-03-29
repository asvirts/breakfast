package main

import "fmt"

func main()  {
	whatWasSaid, whatAge := say()
	fmt.Println(whatWasSaid, whatAge)
}

func say() (string, int) {
	return "Hello, I am", 33
}