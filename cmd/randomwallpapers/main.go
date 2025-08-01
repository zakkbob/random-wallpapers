package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zakkbob/random-wallpapers/internal"
)

// Used to parse custom seed flag
type seed struct {
	X, Y    int
	R, G, B int
}

type seeds []seed

func (s *seeds) String() string {
	str := strings.Builder{}
	for i, d := range *s {
		str.WriteString(fmt.Sprintf("Seed %d - Position %d, %d; Colour %d, %d, %d\n", i, d.X, d.Y, d.R, d.G, d.B))
	}
	return str.String()
}

func (s *seeds) Set(value string) error {
	v := strings.Split(value, ",")
	if len(v) != 5 {
		return errors.New("(must follow 'x,y,r,g,b' e.g '100,100,120,50,8')")
	}

	msg := "invalid %s value '%s' (must be an integer between 0 and 255)"

	x, err := strconv.Atoi(v[0])
	if err != nil {
		return fmt.Errorf(msg, "x", v[0])
	}

	y, err := strconv.Atoi(v[1])
	if err != nil {
		return fmt.Errorf(msg, "y", v[1])
	}

	r, err := strconv.Atoi(v[2])
	if err != nil {
		return fmt.Errorf(msg, "r", v[2])
	}

	g, err := strconv.Atoi(v[3])
	if err != nil {
		return fmt.Errorf(msg, "g", v[3])
	}

	b, err := strconv.Atoi(v[4])
	if err != nil {
		return fmt.Errorf(msg, "b", v[4])
	}

	*s = append(*s, seed{
		X: x,
		Y: y,
		R: r,
		G: g,
		B: b,
	})
	return nil
}

func (s *seeds) add(f *internal.FloodFill) {
	for _, d := range *s {
		f.NewSeed(d.X, d.Y, d.R, d.G, d.B)
	}
}

var seedsFlag seeds
var width int
var height int
var monitor string
var output string
var redMul float64
var greenMul float64
var blueMul float64

func init() {
	defaultPath := filepath.Join(os.TempDir(), "random-wallpaper.png")
	flag.StringVar(&output, "output", defaultPath, "Image save location (including filename and .png extension)")
	flag.IntVar(&width, "width", 1000, "Width of the image")
	flag.IntVar(&height, "height", 1000, "Height of the image")
	flag.StringVar(&monitor, "monitor", "", "Name of monitor to apply the wallpaper to (Run hyprctl monitors to list them)")
	flag.Float64Var(&redMul, "rv", 1, "Red variability")
	flag.Float64Var(&greenMul, "gv", 1, "Green variability")
	flag.Float64Var(&blueMul, "bv", 1, "Blue variability")
	flag.Var(&seedsFlag, "seed", "Add seeds (Usage: --seed x,y,r,g,b e.g --seed 100,100,200,50,80) (Can be used multiple times)")
}

func main() {
	flag.Parse()

	fmt.Print(seedsFlag.String())

	redMul = internal.Clamp(redMul, 0, 255)
	greenMul = internal.Clamp(greenMul, 0, 255)
	blueMul = internal.Clamp(blueMul, 0, 255)

	f := internal.NewFloodFill(width, height)
	f.SetMul(redMul, greenMul, blueMul)

	seedsFlag.add(&f)

	f.Generate()

	err := f.SavePNG(output)
	if err != nil {
		panic(err)
	}

	output, err = filepath.Abs(output)
	if err != nil {
		panic(err)
	}
	if monitor != "" {
		internal.SetWallpaper(monitor, output)
	}
}
