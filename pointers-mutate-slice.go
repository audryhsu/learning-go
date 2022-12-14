package main

import (
	"fmt"
)

func main() {
	lista := []string{"a"}
	fmt.Printf("address of list a: %p\n\n", &lista)

	// pass in a pointer to list a to mutate the list
	mutatemyslicepls(&lista, "o")
	fmt.Println(lista)
}

// KEY: slices are passed by value, like everything else in Go.
// s is a new local variable that is a pointer to a slice of strings
func mutatemyslicepls(s *[]string, letter string) {
	fmt.Printf("memory address of variable s %p\n", &s)
	// KEY: do not confused memory address of s itself vs value of s, which is also a memory address/pointer
	fmt.Printf("value of s holds a pointer to lista: %p\n", s) // s points to same address as lista, just saved in a different place
	fmt.Printf("dereference s: %v\n", *s)

	// append a letter to the slice that s points to
	*s = append(*s, letter)
	fmt.Printf("address of s after append: %p\n", s)
}