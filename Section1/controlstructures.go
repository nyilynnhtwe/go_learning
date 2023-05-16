package main

import "fmt"

func main() {
	age := 22
	if age < 18 {
		fmt.Println("You are not adult")
	} else if age == 18 {
		fmt.Println("You are eighteen")
	} else {
		fmt.Println("You are adult")
	}

	end := 10
	for start := 0; start < end; start++ {
		fmt.Println(start)
	}
}
