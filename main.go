package main

import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/color"
	"math"
)

const FILE_ONE_NAME = "mustang.jpeg"
const FILE_SAME = "mustang_same.jpeg"
const FILE_CHANGED = "mustang_changed.jpeg"

func main() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	file1, err := os.Open(fmt.Sprintf("/Users/wls/code/go-lang/chenghuan/%s", FILE_ONE_NAME))
	file2, err := os.Open(fmt.Sprintf("/Users/wls/code/go-lang/chenghuan/%s", FILE_CHANGED))
	if err != nil {
		fmt.Println("Cannot open image")
		os.Exit(1)
	}

	defer file1.Close()

	img1, _, err := image.Decode(file1)
	img2, _, err := image.Decode(file2)

	rgbaImg1 := convertRGBA(img1)
	rgbaImg2 := convertRGBA(img2)
	o, err := os.Create("/Users/wls/code/go-lang/chenghuan/new_mustang.jpeg")

	fmt.Println("difference is: %d",   compare(rgbaImg1, rgbaImg2))

	if err != nil {
		fmt.Println("Cannot create image.")
		os.Exit(1)
	}
	defer o.Close()
	jpeg.Encode(o, rgbaImg1, nil)
}

func compare(img1, img2 *image.RGBA) int64 {
	diff := int64(0)
	for i := 0; i < len(img1.Pix); i++ {
		diff += int64(diffPixel(img1.Pix[i], img2.Pix[i]))
	}
	return int64(math.Sqrt(float64(diff)))
}

func diffPixel(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

func convertRGBA(img image.Image) (*image.RGBA){
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	for y:=0; y < bounds.Max.Y; y++ {
		for x:=0; x < bounds.Max.X; x++ {
			//old := img.At(x,y)
			//r, g, b, a := old.RGBA()
			//pixel := color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
			rgbaImg.Set(x, y, color.RGBAModel.Convert(img.At(x, y)))
		}
	}
	return rgbaImg
}