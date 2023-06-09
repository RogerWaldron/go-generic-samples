package main

import "fmt"

func GetOrDefault[K comparable, V any](m map[K]V, key K, useDefault func() V) V {
	value, found := m[key]

	if !found {
		value = useDefault()
		m[key] = value
	}

	return value
}

func main() {
	a := make(map[string]string)
	a["Server_Port"] = "4200"
	fmt.Println(GetOrDefault(a, "Server_Name", func() string { return "localhost"})) // localhost
	fmt.Println(GetOrDefault(a, "Server_Port", func() string { return "3000"})) // 4200

	b := make(map[string]int)
	b["GET"] = 200
	b["PUT_Created"] = 201
	b["PUT_Updated"] = 204
	b["DELETE"] = 204
	fmt.Println(GetOrDefault(b, "PUT_Created", func() int { return 200 })) // 201
	fmt.Println(GetOrDefault(b, "POST", func() int { return 200 })) // 200
}