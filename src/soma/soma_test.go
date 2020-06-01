package main

import "testing"

func TestSomaUmMaisDois(t *testing.T) {
	x := soma(1, 2)
	if x != 3 {
		t.Error("Opa! 1+2 não é igual a 3, obtive", x)
	}
}

func TestSomaUmMaisUm(t *testing.T) {
	x := soma(1, 1)
	if x != 2 {
		t.Error("Opa! 1+1 não é igual a 2, obtive", x)
	}
}
