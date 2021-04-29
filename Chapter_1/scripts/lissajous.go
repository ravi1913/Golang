//lissajous generates GIF animations of random lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0x4f, 0xeb, 0x34, 0xff},
	color.RGBA{0xeb, 0x49, 0x34, 0xff},
	color.RGBA{0xeb, 0xe1, 0x34, 0xff},
	color.RGBA{0x34, 0xeb, 0xcc, 0xff},
	color.RGBA{0x34, 0xc6, 0xeb, 0xff},
	color.RGBA{0x34, 0x3a, 0xeb, 0xff},
	color.RGBA{0x89, 0x34, 0xeb, 0xff},
	color.RGBA{0xeb, 0x34, 0xcf, 0xff},
}

const (
	whiteIndex  = 0
	blackIndex  = 1
	greenIndex  = 2
	redIndex    = 3
	yellowIndex = 4
	cyanIndex   = 5
	blueIndex   = 6
	indigoIndex = 7
	purpleIndex = 8
	pinkIndex   = 9
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 1 + rand.Intn(9-0+1)
	newIndex := uint8(n)
	lissajous(os.Stdout, newIndex)
}
func lissajous(out io.Writer, index uint8) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
