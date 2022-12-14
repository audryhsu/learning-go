package main

import (
	"fmt"
)

// Two ways of mutating an array from another function
// 1. Pass a pointer to array to the mutating function
// 2. Pass copy of slice to function, reassign original array to the returned slice
func main() {
	lista := []string{"a"}
	fmt.Printf("address of list a: %p\n\n", &lista)
	originalAddress := &lista

	// pass in a pointer to lista to mutate the list
	mutate1(&lista, "o")
	fmt.Println("List a contents: ", lista)

	// pass a copy of lista and overwrite slice in same memory address
	lista = mutate2(lista)

	fmt.Println("mutate2 mutates lista:", lista)
	res := originalAddress == &lista
	fmt.Printf("original address === &lista: %v", res)
}

// Slices are pass by value. s is a copy of a slice of strings
// mutate2 appends a letter to slice s and returns it.
func mutate2(s []string) []string {
	fmt.Printf("address of s: %p\n", &s)
	s = append(s, "o")
	return s
}

// KEY: slices are passed by value, like everything else in Go.
// s is a new local variable that is a pointer to a slice of strings
func mutate1(s *[]string, letter string) {
	// KEY: do not confused memory address of s itself vs value of s, which is also a memory address/pointer
	fmt.Printf("memory address of variable s %p\n", &s)
	fmt.Printf("value of s holds a pointer to lista: %p\n", s) // s points to same address as lista, just saved in a different place
	fmt.Printf("dereference s gets the value that s points to: %v\n", *s)

	// append a letter to the slice that s points to
	*s = append(*s, letter)
}