package internal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
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

func pow(x, y int) int {
	t := x
	for range y - 1 {
		x *= t
	}
	return x
}

func dist(p0, p1 *image.Point) float64 {
	d := math.Sqrt(float64(pow(p0.X-p1.X, 2) + pow(p0.Y-p1.Y, 2)))
	return d
}

func NewImageFloodFill(m image.Image, r int) FloodFill {
	f := NewFloodFill(m.Bounds().Dx(), m.Bounds().Dy())

	width := m.Bounds().Dx()
	height := m.Bounds().Dy()

	active := make([]*image.Point, 0)

	cellSize := math.Floor(float64(r) / math.Sqrt2)

	xCells := int(math.Ceil(float64(width)/cellSize)) + 1
	yCells := int(math.Ceil(float64(height)/cellSize)) + 1

	grid := make([][]*image.Point, xCells)
	for i := range xCells {
		grid[i] = make([]*image.Point, yCells)
	}

	addSeed := func(p *image.Point) {
		x := int(math.Floor(float64(p.X) / cellSize))
		y := int(math.Floor(float64(p.Y) / cellSize))
		grid[x][y] = p

		r, g, b, _ := m.At(x, y).RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		f.NewSeed(x, y, int(r), int(g), int(b))
	}

	validPoint := func(p *image.Point) bool {
		if p.X < 0 || p.X >= width || p.Y < 0 || p.Y >= height {
			return false
		}

		x := int(math.Floor(float64(p.X) / cellSize))
		y := int(math.Floor(float64(p.Y) / cellSize))
		i0 := max(x-1, 0)
		i1 := min(x+1, xCells-1)
		j0 := max(y-1, 0)
		j1 := min(y+1, yCells-1)

		for i := i0; i <= i1; i++ {
			for j := j0; j <= j1; j++ {
				if grid[i][j] != nil && dist(grid[i][j], p) < float64(r) {
					return false
				}
			}
		}

		return true
	}

	initSeed := &image.Point{
		X: rand.Intn(width),
		Y: rand.Intn(height),
	}
	active = append(active, initSeed)
	addSeed(initSeed)

	for len(active) > 0 {
		i := rand.Intn(len(active))
		p := active[i]

		found := false
		for range 30 {
			theta := rand.Float64() * math.Pi * 2
			newRadius := rand.Float64()*float64(r) + float64(r)
			newP := &image.Point{
				X: p.X + int(newRadius*math.Cos(theta)),
				Y: p.Y + int(newRadius*math.Sin(theta)),
			}

			if !validPoint(newP) {
				continue
			}

			addSeed(p)
			active = append(active, newP)
			found = true
			break
		}

		if !found {
			active[i] = active[len(active)-1]
			active[len(active)-1] = nil
			active = active[:len(active)-1]
		}
	}

	return f
}

func (f *FloodFill) NewSeed(x, y int, r, g, b int) {
	c := color.RGBA{}

	c.A = 0
	c.R = uint8(Clamp(r, 0, 255))
	c.G = uint8(Clamp(g, 0, 255))
	c.B = uint8(Clamp(b, 0, 255))

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
	c.G = uint8(Clamp(int(c.G)+int(rand.NormFloat64()*f.gMul), 0, 255))
	c.B = uint8(Clamp(int(c.B)+int(rand.NormFloat64()*f.bMul), 0, 255))
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
