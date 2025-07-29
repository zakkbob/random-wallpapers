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
	monitor := flag.String("monitor", "", "Name of monitor to apply the wallpaper to (Run hyprctl monitors to list them)")
	cR := flag.Int("sr", 128, "Seed colour (Red channel)")
	cG := flag.Int("sg", 128, "Seed colour (Green channel)")
	cB := flag.Int("sb", 128, "Seed colour (Blue channel)")
	rV := flag.Float64("rv", 1, "Red variability")
	gV := flag.Float64("gv", 1, "Green variability")
	bV := flag.Float64("bv", 1, "Blue variability")

	flag.Parse()

	if *monitor == "" {
		fmt.Print("--monitor flag is required")
		os.Exit(1)
	}

	*cR = internal.Clamp(*cR, 0, 255)
	*cG = internal.Clamp(*cG, 0, 255)
	*cB = internal.Clamp(*cB, 0, 255)

	f := internal.NewFloodFill(*width, *height)
	f.NewSeed(*width/2, *height/2, *cR, *cG, *cB)
	f.SetMul(*rV, *gV, *bV)
	f.Generate()

	dir := os.TempDir() + "/dynamic-wallpaper-image.png"

	err := f.SavePNG(dir)
	if err != nil {
		panic(err)
	}

	internal.SetWallpaper(*monitor, dir)
}
