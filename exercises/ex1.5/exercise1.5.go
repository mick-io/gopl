// Exercise 1.5: Change the Lissajous program's color pallet to green on black. for added authenticity, to create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair of hexadecimal digits represent the intensity of red, green, or blue componets of pixels.
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
	green   = color.RGBA{0, 255, 0, 100}
	palette = []color.Color{color.Black, green}
)

const (
	blackIndex = 0 // first color in pallet
	greenIndex = 1 // next color in pallet
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
		img := image.NewPaletted(rect, palette) //

		// Each pass through the inner loop generates a new image by setting some pixels to black.
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// The inner loop runs the two oscillators. THe x oscillator is just the sine function. THe y oscillator
			// is also a sinusoid, but its frequency relative to the x oscillator is random number between 0 and 3, and
			// its phase relative to the x oscillator is initially zero but increase with each frame of the animation.
			// The loop runs until until the x oscillator has completed 5 cycles. At each step, it calls SetColorIndex
			// to color the pixel corresponding to (x,y) black, which is at position 1 in the pallet.
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
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
