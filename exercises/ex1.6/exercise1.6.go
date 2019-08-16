// Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding more values to palette
// and then displaying them by changing the third argument of SetColorIndex in some interesting way.
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

var (
	red     = color.RGBA{255, 0, 0, 100}
	green   = color.RGBA{0, 255, 0, 100}
	blue    = color.RGBA{0, 0, 255, 100}
	palette = []color.Color{color.Black, red, blue, green}
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// The outer loop runs for 64 iterations, each producing a single frame of the animation.
	for i := 0; i < nframes; i++ {
		// Creates a new 201x201 image with the pallet of two colors.
		// Pixels are initially set to the pallet's zero value (the zeroth color in the palette).
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Each pass through the inner loop generates a new image by changing the color of some pixels.
		var colorIndex uint8 = 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
			colorIndex++
			if colorIndex > 3 {
				colorIndex = 1
			}
		}
		phase += 0.1
		// The results are appended using the built-in append function
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		panic(err)
	}
}
