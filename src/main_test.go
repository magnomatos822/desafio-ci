package main

import (
	"modulo/somar"
	"testing"
)

func TestSum(t *testing.T) {
	resultado := somar.Somar(5, 5)

	if resultado != 10 {
		t.Errorf("Soma retorn = %d; correto 10", resultado)
	}
}
