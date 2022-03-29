package identigon

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/image/draw"
)

func TestGenerate(t *testing.T) {
	img := Generate("seed", 4, 4)
	assert.Equal(t, [][][3]float32{
		{{124, 242, 229}, {124, 242, 229}, {124, 242, 229}, {124, 242, 229}},
		{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		{{124, 242, 229}, {124, 242, 229}, {124, 242, 229}, {124, 242, 229}},
		{{124, 242, 229}, {0, 0, 0}, {0, 0, 0}, {124, 242, 229}},
	}, image_2_array_pix(img))

}

func image_2_array_pix(src image.Image) [][][3]float32 {
	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	iaa := make([][][3]float32, height)
	src_rgba := image.NewRGBA(src.Bounds())
	draw.Copy(src_rgba, image.Point{}, src, src.Bounds(), draw.Src, nil)
	for y := 0; y < height; y++ {
		row := make([][3]float32, width)
		for x := 0; x < width; x++ {
			idx_s := (y*width + x) * 4
			pix := src_rgba.Pix[idx_s : idx_s+4]
			row[x] = [3]float32{float32(pix[0]), float32(pix[1]), float32(pix[2])}
		}
		iaa[y] = row
	}
	return iaa
}
