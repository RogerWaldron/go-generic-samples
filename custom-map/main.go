package main

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V] {
		data: make(map[K]V),
	}
}
func main() {
	map1 := NewCustomMap[string, int]()
	map1.Insert("Sydney", 2000)
	map1.Insert("Salt Lake City", 2002)
	map1.Insert("Athens", 2004)
	map1.Insert("Turin", 2006)
}