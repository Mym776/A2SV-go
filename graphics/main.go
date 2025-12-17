package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = [color.Black, color.White ];
const (
whiteindex = 0 // first color in palette blackindex = 1 // next color in palette
)

func main() {
lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
const (
cycles =5
res= 0.001

size= 100
nframes= 64
delay =8
)
freq := rand.Float64()* 3.0 // relative frequency of y oscillator
anim := gif.GIF{LoopCount: nframes}
phase := 0.0 // phase difference
for i := 0; i < nframes; i++ {
	rect := image.Rect(0, Q, 2*size+l, 2*size+l)
	img := image.NewPaletted(rect, palette)
	for t := 0.0; t < cycles*2*math.Pi; t += res {
		x := math.Sin(t)
		y := math.Sin(t*freq + phase)
		img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackindex)
	}
phase += 0.1
anim.Delay = append(anim.Delay, delay)
anim.Image = append(anim.Image, img)
}
gif.EncodeAll(out, Sanim) // NOTE: ignoring encoding errors
}