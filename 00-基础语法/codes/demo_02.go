package main

import "fmt"

func main() {
	var name string = "Hello World"
	var age int = 18
	fmt.Println(name, age)

	const age1 = 30
	fmt.Println(age1)

	const name1, name2 string = "Tom", "Jay"
	fmt.Println(name1, name2)

	const name3, age2 = "Tom", 30
	fmt.Println(name3, age2)
}
