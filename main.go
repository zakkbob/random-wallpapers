package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func changeWallpaper(dir string) error {
	monitor := "eDP-1"
	arg := monitor + "," + dir
	_, err := exec.Command("hyprctl", "hyprpaper", "wallpaper", arg).Output()
	if err != nil {
		return fmt.Errorf("failed to change wallpaper: %w", err)
	}
	return nil
}

func savePNG(m image.Image, name string) error {
	f, err := os.Create(name + ".png")
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	err = png.Encode(f, m)
	if err != nil {
		return fmt.Errorf("encoding and saving image: %w", err)
	}

	return nil
}

func clamp(n int, u int, l int) int {
	if n < l {
		return l
	}
	if n > u {
		return u
	}
	return n
}

func perturb(m *image.RGBA, seeds []image.Point, seed image.Point, p image.Point) []image.Point {
	r, g, b, a := m.At(p.X, p.Y).RGBA()
	if r != 0 || g != 0 || b != 0 || a != 0 { // if the pixel has been changed already
		return seeds
	}

	seeds = append(seeds, p)

	r, g, b, a = m.At(seed.X, seed.Y).RGBA()

	// do a random Guassian perturbation on all channels
	c := color.RGBA{}
	c.A = 255

	c.R = uint8(r >> 8)
	c.G = uint8(g >> 8)
	c.B = uint8(b >> 8)

	var mul float64 = 5

	var (
		rR = int(rand.NormFloat64() * mul)
		rG = int(rand.NormFloat64() * mul)
		rB = int(rand.NormFloat64() * mul)
	)

	c.R = uint8(clamp(int(c.R)+rR, 255, 0))
	c.G = uint8(clamp(int(c.G)+rG, 255, 0))
	c.B = uint8(clamp(int(c.B)+rB, 255, 0))

	m.Set(p.X, p.Y, c)
	return seeds
}

func newFloodFill(width int, height int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	initialSeed := image.Point{X: width / 2, Y: height / 2}
	m.Set(initialSeed.X, initialSeed.Y, color.RGBA{R: 100, G: 70, B: 28, A: 255})

	seeds := make([]image.Point, 0)
	seeds = append(seeds, initialSeed)

	for range 10000000 {
		seed := seeds[rand.Intn(len(seeds))]

		seeds = perturb(m, seeds, seed, seed.Add(image.Point{-1, -1}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{-1, 0}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{-1, 1}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{0, -1}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{0, 1}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{1, -1}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{1, 0}))
		seeds = perturb(m, seeds, seed, seed.Add(image.Point{1, 1}))
	}

	return m
}

func main() {
	m := newFloodFill(1000, 1000)

	err := savePNG(m, "image")
	if err != nil {
		panic(err)
	}
}

func hyprpaperChange() {
	FPS := 10
	var err error
	for {
		time.Sleep(time.Millisecond * time.Duration(1000/FPS))
		_, err = exec.Command("hyprctl", "hyprpaper", "wallpaper", "eDP-1,~/Downloads/a_moon_over_a_mountain.png").Output()
		if err != nil {
			fmt.Printf("%v", err)
		}
	}

}
