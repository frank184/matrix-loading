package main

import "testing"

func BenchmarkRandomSymbol(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomSymbol()
	}
}

func BenchmarkSymbols(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Symbols()
	}
}
