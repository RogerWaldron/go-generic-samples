package main

import "fmt"

func coalesce[T comparable](values ...T) (zero T) {
	for _, item := range values {
		if item != zero {
			return item
		}
	}
	
	return zero
}

func main() {
	// return the first non zero value
	fmt.Println(coalesce("a", "b")) // "a"
	fmt.Println(coalesce("", "b"))	// "b"
	fmt.Println(coalesce(0, 0, 200)) // 200
	fmt.Println(coalesce(false, true))	// true
}