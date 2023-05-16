package main

import "fmt"

func main() {
	var age int
	var name string
	var weight float32
	var isAdult bool

	age = 22
	name = "david"
	weight = 150.55
	isAdult = true

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(weight)
	fmt.Println(isAdult)

	// rune example
	str := "Thank you <-> ありがと"
	for _, val := range str {
		fmt.Printf("unicode : %U val : %c\n", val, val)
	}

	// byte example
	var word byte
	word = 65
	fmt.Printf("byte code : %d character : %c\n", word, word)

}
