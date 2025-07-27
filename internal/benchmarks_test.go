package internal_test

import (
	"testing"

	"github.com/zakkbob/dynamic-wallpaper/internal"
)

func BenchmarkFloodFill100(b *testing.B) {
	for b.Loop() {
		internal.NewFloodFill(100, 100)
	}
}

func BenchmarkFloodFill500(b *testing.B) {
	for b.Loop() {
		internal.NewFloodFill(500, 500)
	}
}
