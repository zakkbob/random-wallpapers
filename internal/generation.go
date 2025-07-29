package internal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

type FloodFill struct {
	width  int
	height int

	image *image.RGBA
	seeds []image.Point

	rMul float64
	gMul float64
	bMul float64

	generated bool
}

func NewFloodFill(width int, height int) FloodFill {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	return FloodFill{
		width:     width,
		height:    height,
		image:     m,
		rMul:      1,
		gMul:      1,
		bMul:      1,
		generated: false,
	}
}

func (f *FloodFill) NewSeed(x int, y int, r int, g int, b int) {
	c := color.RGBA{}

	c.A = 0
	c.R = uint8(Clamp(r, 0, 255))
	c.G = uint8(Clamp(r, 0, 255))
	c.B = uint8(Clamp(r, 0, 255))

	f.seeds = append(f.seeds, image.Point{X: x, Y: y})
	f.image.SetRGBA(x, y, c)
}

func (f *FloodFill) SetMul(rMul float64, gMul float64, bMul float64) {
	f.rMul = rMul
	f.gMul = gMul
	f.bMul = bMul
}

func (f *FloodFill) randSeed() (int, image.Point) {
	i := rand.Intn(len(f.seeds))
	return i, f.seeds[i]
}

func (f *FloodFill) changed(p image.Point) bool {
	c := f.image.RGBAAt(p.X, p.Y)
	return c.A == 255
}

func (f *FloodFill) grow(seed image.Point, dX int, dY int) {
	p := image.Point{
		X: seed.X + dX,
		Y: seed.Y + dY,
	}

	if p.X < 0 || p.X >= f.width || p.Y < 0 || p.Y >= f.height {
		return
	}

	if f.changed(p) { // if the pixel has been changed already
		return
	}

	f.seeds = append(f.seeds, p)

	c := f.image.RGBAAt(seed.X, seed.Y)

	c.R = uint8(Clamp(int(c.R)+int(rand.NormFloat64()*f.rMul), 0, 255))
	c.G = uint8(Clamp(int(c.G)+int(rand.NormFloat64()*f.rMul), 0, 255))
	c.B = uint8(Clamp(int(c.B)+int(rand.NormFloat64()*f.rMul), 0, 255))
	c.A = 255

	f.image.SetRGBA(p.X, p.Y, c)
}

func (f *FloodFill) removeSeed(i int) {
	f.seeds[i] = f.seeds[len(f.seeds)-1]
	f.seeds = f.seeds[:len(f.seeds)-1]
}

func (f *FloodFill) Generate() {
	if f.generated {
		panic("flood fill has already been generated!")
	}
	f.generated = true
	for {
		if len(f.seeds) == 0 {
			break
		}

		i, seed := f.randSeed()

		f.grow(seed, -1, -1)
		f.grow(seed, -1, 0)
		f.grow(seed, -1, 1)
		f.grow(seed, 0, -1)
		f.grow(seed, 0, 1)
		f.grow(seed, 1, -1)
		f.grow(seed, 1, 0)
		f.grow(seed, 1, 1)

		f.removeSeed(i)
	}
}

func (f *FloodFill) Image() *image.RGBA {
	return f.image
}

func (f *FloodFill) SavePNG(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	err = png.Encode(file, f.image)
	if err != nil {
		return fmt.Errorf("encoding and saving image: %w", err)
	}

	return nil
}
