package rnd

import (
	"testing"
)

func TestRandom1(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Random(0, 1))
	}
}
func TestRandom2(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Random(-2, -1))
	}
}

func TestRandom3(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Random(-2, 2))
	}
}
