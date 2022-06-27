package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(numbers ...int) int {
	var result int
	for _, num := range numbers {
		result = result + num
	}

	return result
}
