package main

import (
	"fmt"
	"image"
	"secretmerger/internal/imageio"
	"secretmerger/internal/process"
)

func main() {
	// x := byte(54) // 00110110
	// y := process.ReverseByte(x)
	// z := process.ReverseBits(x)
	// fmt.Printf("Original: %08b\n", x) // 輸出原始位元
	// fmt.Printf("Reversed: %08b\n", y) // 輸出翻轉後的位元
	// fmt.Printf("Reversed Bits: %08b\n", z) // 輸出翻轉後的位元(只保留後四位並往左移)

	inputX := "testdata/baboon_gray.bmp"
	headerX, pixelsX, err := imageio.ReadGrayBMP(inputX)
	if err != nil {
		panic(err)
	}
	widthX := int(headerX[18]) | int(headerX[19])<<8 | int(headerX[20])<<16 | int(headerX[21])<<24
	heightX := int(headerX[22]) | int(headerX[23])<<8 | int(headerX[24])<<16 | int(headerX[25])<<24
	// 將像素轉換為 *image.Gray
	imgX := &image.Gray{
		Pix:    pixelsX,
		Stride: widthX,
		Rect:   image.Rect(0, 0, widthX, heightX),
	}

	outputXReversed := "output/outputXReversed.bmp"
	
	fmt.Println("Before: ", imgX.Pix[:16])
	reversedX := process.ReverseImageBits(imgX)
	if err := imageio.WriteGrayBMP(outputXReversed, headerX, reversedX.Pix, widthX, heightX); err != nil {
		panic(err)
	}
	fmt.Println("After: ", reversedX.Pix[:16])
}
