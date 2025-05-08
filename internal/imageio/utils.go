package imageio

import (
	"image"
)

func IsGrayscale(img image.Image) bool {
    _, ok := img.(*image.Gray)
    return ok
}

func SameSize(img1, img2 image.Image) bool {
    return img1.Bounds().Size() == img2.Bounds().Size()
}