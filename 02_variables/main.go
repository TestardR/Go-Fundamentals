package main

import "fmt"

func main() {
	name := "Romain"
	size := 1.3
	email, server := "romain.rtestard@gmail.com", "myserver"
	var age int32 = 32
	const isCool bool = true

	fmt.Println(name, age, isCool, size, email, server)
	fmt.Printf("%T", isCool)
}
