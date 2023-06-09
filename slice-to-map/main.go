package main

import "fmt"

type Thing struct {
	Value string
}

func Index[K comparable, V any](list []V, key func(V) K) map[K]V {
	result := make(map[K]V)

	for _, value := range list {
		result[key(value)] = value
	}

	return result
}

func main() {
	// Index a slice of Objects into a Map
	a := []Thing{{Value: "Hello"}, {Value: "World"}}
	b := Index(a, func(t Thing) string { return t.Value })
	fmt.Println(b["Hello"]) // {Hello}
}