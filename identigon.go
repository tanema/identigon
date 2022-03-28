package identigon

import (
	"crypto/sha512"
	"image"
	"image/color"

	"golang.org/x/image/draw"
)

// Generate will create a new identicon image from a string of data
// size: square size of the image in pixels (i.e. 80)
// pixels: how many block should the image be made up of. (4x4 = 8)
// border: thickness of borders in pixels.
func Generate(data string, size, pixels, border int) image.Image {
	pix := size / pixels
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	sum := sha512.Sum512([]byte(data))
	pixcolor := color.RGBA{uint8(sum[0]), uint8(sum[1]), uint8(sum[2]), 255}
	borcolor := color.RGBA{uint8(sum[3]), uint8(sum[4]), uint8(sum[5]), 255}
	i := 0
	for y := 0; y < pixels; y++ {
		for x := 0; x < pixels/2; x++ {
			if sum[i%len(sum)]&1 == 0 {
				drawRect(img, x*pix, y*pix, pix, pix, pixcolor)
				drawRect(img, (pixels-1-x)*pix, y*pix, pix, pix, pixcolor)
			}
			i++
		}
	}
	drawBorder(img, size, border, borcolor)
	return img
}

func drawBorder(img draw.Image, imgSize, borderSize int, clr color.RGBA) {
	drawRect(img, 0, 0, imgSize, borderSize, clr)
	drawRect(img, 0, 0, borderSize, imgSize, clr)
	drawRect(img, imgSize-borderSize, 0, borderSize, imgSize, clr)
	drawRect(img, 0, imgSize-borderSize, imgSize, borderSize, clr)
}

func drawRect(img draw.Image, x, y, w, h int, clr color.RGBA) {
	draw.Draw(img, image.Rect(x, y, x+w, y+h), &image.Uniform{clr}, image.ZP, draw.Src)
}
