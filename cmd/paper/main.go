package main

import (
	"os"

	"github.com/zakkbob/dynamic-wallpaper/internal"
)

func main() {
	m := internal.NewFloodFill(900, 450)

	dir := os.TempDir() + "/image.png"

	err := internal.SavePNG(m, dir)
	if err != nil {
		panic(err)
	}

	internal.SetWallpaper(dir)
}
