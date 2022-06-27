package main

import "fmt"

func Alphabet(length int) []string {
	var characters []string
	for i := 0; i < length; i++ {
		characters = append(characters, characterByIndex(i))
	}

	return characters
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
