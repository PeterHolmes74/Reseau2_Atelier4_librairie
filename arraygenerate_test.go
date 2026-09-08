package main

import (
	"testing"
)

func TestArrayGenerate(t *testing.T) {
	attendu := len(ArrayGenerate(-1)) == 0

	if !attendu {
		t.Errorf("ArrayGenerate de mauvaise taille %d", len(ArrayGenerate(-1)))
	}
}

func TestArrayGenerate2(t *testing.T) {
	obtenu := ArrayGenerate(1)

	if !(obtenu[0] >= 1 && obtenu[0] <= 10000) {
		t.Errorf("Contenu de ArrayGenerate de mauvaise taille %d", obtenu[0])
	}
}
