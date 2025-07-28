package internal_test

import (
	"image/color"
	"testing"

	"github.com/zakkbob/dynamic-wallpaper/internal"
)

func BenchmarkFloodFill100(b *testing.B) {
	for b.Loop() {
		internal.NewFloodFill(100, 100, color.RGBA{R: 128, G: 128, B:  128}, 1, 1, 1)
	}
}

func BenchmarkFloodFill500(b *testing.B) {
	for b.Loop() {
		internal.NewFloodFill(500, 500, color.RGBA{R: 128, G: 128, B:  128}, 1, 1, 1)
	}
}
