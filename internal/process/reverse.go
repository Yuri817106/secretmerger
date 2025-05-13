package process

import (
	"image"
	"image/color"
)

// ReverseByte 反轉一個 byte 的 bit 順序
func ReverseByte(b byte) byte {
	var r byte = 0
	for i := 0; i < 8; i++ {
		r <<= 1 // r 左移
		r |= b & 1 // 取 b 最右邊的 bit (b & 1) 放到 r 最右邊的 bit (r |= ...) 
		b >>= 1 // b 右移
	}
	return r
}
// ReverseBits 反轉一個 byte 的 bit 順序，並保存成 B4B5B6B7 0000
func ReverseBits(b byte) byte {
	var r byte = 0
	for i:= 0; i < 8; i++ {
		r <<= 1
		r |= b & 1
		b >>= 1
	}
	r <<= 4
	return r
}

// ReverseImageBits 對每個像素的 bit 順序做反轉（僅支援灰階圖）
func ReverseImageBits(img *image.Gray) *image.Gray {
	bounds := img.Bounds()
	newImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			orig := img.GrayAt(x, y).Y
			newVal := ReverseByte(orig)
			newImg.SetGray(x, y, color.Gray{Y: newVal})
		}
	}
	return newImg
}
