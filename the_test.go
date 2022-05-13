package main

import "testing"

func TestDecimal(t *testing.T) {
	if decimal() != 1539 {
		t.Error()
	}
}

func TestBinary(t *testing.T) {
	if binary() != 1539 {
		t.Error()
	}
}

func BenchmarkDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimal()
	}
}

func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binary()
	}
}
