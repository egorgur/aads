package main

import (
	"fmt"
	"pr3/internal"
)

func main() {

	// Пример использования
	text := "abc"
	subStrings := []string{"abc"}
	caseSensitivity := false
	method := "last"
	count := 1
	result := search.Search(text, subStrings, caseSensitivity, method, count)
	fmt.Println(result)

}
