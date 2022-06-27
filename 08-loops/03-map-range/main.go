package main

import "fmt"

// MAPS ARE NOT ORDERED!
//Calling range on a map returns the values in a random order, so you can't rely on it.

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(mapa map[int]string) []int {
	var chaves []int
	for k, _ := range mapa {
		chaves = append(chaves, k)
	}
	return chaves
}

func Values(mapa map[int]string) []string {
	var valores []string
	for _, val := range mapa {
		valores = append(valores, val)
	}
	return valores
}
