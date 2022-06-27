package main

import "fmt"

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}

func Swap(str1, str2 string) (string, string) {
	return str2, str1
}
