package main

import "fmt"

func WordGenerator(words []string) func() string {
	counter := 0
	return func() string {
		toReturn := words[counter]
		counter++
		if counter == len(words) {
			counter = 0
		}
		return toReturn
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
