package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zakkbob/dynamic-wallpaper/internal"
)

func main() {
	width := flag.Int("width", 100, "Width of the image")
	height := flag.Int("height", 100, "Height of the image")
	monitor := flag.String("monitor", "", "Name of monitor to apply the wallpaper to (Run `hyprctl monitors` to see them)")

	flag.Parse()

	if *monitor == "" {
		fmt.Print("--monitor flag is required")
		os.Exit(1)
	}

	m := internal.NewFloodFill(*width, *height)

	dir := os.TempDir() + "/dynamic-wallpaper-image.png"

	err := internal.SavePNG(m, dir)
	if err != nil {
		panic(err)
	}

	internal.SetWallpaper(*monitor, dir)
}
