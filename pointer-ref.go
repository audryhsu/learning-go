package main

import "fmt"

func main() {
	a := struct {
		message string
	}{"hello"}

	b := &a

	fmt.Println(b)
}
