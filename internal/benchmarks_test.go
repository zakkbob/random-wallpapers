package internal_test

import (
	"testing"

	"github.com/zakkbob/random-wallpapers/internal"
)

func BenchmarkFloodFill100(b *testing.B) {
	for b.Loop() {
		f := internal.NewFloodFill(100, 100)
		f.NewSeed(50, 50, 128, 128, 128)
		f.SetMul(1, 1, 1)
		f.Generate()
	}
}

func BenchmarkFloodFill500(b *testing.B) {
	for b.Loop() {
		f := internal.NewFloodFill(500, 500)
		f.NewSeed(250, 250, 128, 128, 128)
		f.SetMul(1, 1, 1)
		f.Generate()
	}
}
