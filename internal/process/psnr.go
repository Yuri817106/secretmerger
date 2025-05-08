package process

import (
	"image"
	"math"
)

var referenceImage *image.Gray // 全域變數，用於存放參考影像

// SetReferenceImage 設定參考影像
func SetReferenceImage(img *image.Gray) {
	referenceImage = img
}

// PSNR 計算輸入影像與參考影像的 PSNR 值
func PSNR(img *image.Gray) float64 {
	if referenceImage == nil {
		panic("reference image is not set")
	}
	if img.Bounds() != referenceImage.Bounds() {
		panic("images must have the same dimensions")
	}

	var mse float64
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			diff := float64(img.GrayAt(x, y).Y) - float64(referenceImage.GrayAt(x, y).Y)
			mse += diff * diff
		}
	}

	mse /= float64(bounds.Dx() * bounds.Dy())
	if mse == 0 {
		return math.Inf(1) // 無限大，表示完全相同
	}

	return 10 * math.Log10(255*255/mse)
}
