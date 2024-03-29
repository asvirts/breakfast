package main

import "fmt"

func main()  {
	whatWasSaid := say()
	whatAge := age()
	fmt.Println(whatWasSaid, whatAge)
}

func say() string {
	return "Hello, I am"
}

func age() int {
	return 33
}