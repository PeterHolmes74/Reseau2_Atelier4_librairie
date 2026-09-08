package main

import "math/rand"

// ArrayGenerate returns a slice containing taille random integers.
func ArrayGenerate(taille int) []int {
	if taille <= 0 {
		return []int{}
	}
	nombres := make([]int, taille)

	for i := range nombres {
		nombres[i] = (rand.Int() % 9999) + 1
	}

	return nombres
}
