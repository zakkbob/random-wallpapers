package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"syscall/js"

	"github.com/zakkbob/random-wallpapers/internal"
)

var floodFill internal.FloodFill

func NewFill(this js.Value, args []js.Value) interface{} {
	width := args[0].Int()
	height := args[1].Int()
	floodFill = internal.NewFloodFill(width, height)
	return nil
}

func AddSeed(this js.Value, args []js.Value) interface{} {
	x := args[0].Int()
	y := args[1].Int()
	r := args[2].Int()
	g := args[3].Int()
	b := args[4].Int()
	floodFill.NewSeed(x, y, r, g, b)
	return nil
}

func Show(this js.Value, args []js.Value) interface{} {
	floodFill.Generate()

	var buf bytes.Buffer
	png.Encode(&buf, floodFill.Image())

	img := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println(img)

	js.Global().Set("img", js.ValueOf(img))
	js.Global().Get("document").Call("getElementById", "img").Set("src", js.Global().Get("img"))
	return nil
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("newFill", js.FuncOf(NewFill))
	js.Global().Set("addSeed", js.FuncOf(AddSeed))
	js.Global().Set("show", js.FuncOf(Show))

	fmt.Println("Hello, WebAssembly!")

	select {}
}
