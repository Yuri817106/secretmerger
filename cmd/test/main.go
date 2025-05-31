package main

import (
	"fmt"
	"image"
	"secretmerger/internal/imageio"
	"secretmerger/internal/process"
)

func main() {
	inputY := "testdata/peppers_gray.bmp"
	inputX := "testdata/baboon_gray.bmp"
	outputX := "output/outputX.bmp"
	outputY := "output/outputY.bmp"
	outputXReversed := "output/outputXReversed.bmp"
	outputYReversed := "output/outputYReversed.bmp"

	// 讀取 X 和 Y 影像
	headerX, pixelsX, err := imageio.ReadGrayBMP(inputX)
	if err != nil {
		panic(err)
	}
	headerY, pixelsY, err := imageio.ReadGrayBMP(inputY)
	if err != nil {
		panic(err)
	}

	// 確保兩張影像大小相同
	widthX := int(headerX[18]) | int(headerX[19])<<8 | int(headerX[20])<<16 | int(headerX[21])<<24
	heightX := int(headerX[22]) | int(headerX[23])<<8 | int(headerX[24])<<16 | int(headerX[25])<<24
	widthY := int(headerY[18]) | int(headerY[19])<<8 | int(headerY[20])<<16 | int(headerY[21])<<24
	heightY := int(headerY[22]) | int(headerY[23])<<8 | int(headerY[24])<<16 | int(headerY[25])<<24

	if widthX != widthY || heightX != heightY {
		panic("Images X and Y must have the same dimensions")
	}

	// 將像素轉換為 *image.Gray
	imgX := &image.Gray{
		Pix:    pixelsX,
		Stride: widthX,
		Rect:   image.Rect(0, 0, widthX, heightX),
	}
	imgY := &image.Gray{
		Pix:    pixelsY,
		Stride: widthY,
		Rect:   image.Rect(0, 0, widthY, heightY),
	}
	
	if err := imageio.WriteGrayBMP(outputX, headerX, imgX.Pix, widthX, heightX); err != nil {
		panic(err)
	}
	if err := imageio.WriteGrayBMP(outputY, headerX, imgY.Pix, widthX, heightX); err != nil {
		panic(err)
	}
	reversedX := process.ReverseImageBits(imgX)
	if err := imageio.WriteGrayBMP(outputXReversed, headerX, reversedX.Pix, widthX, heightX); err != nil {
		panic(err)
	}
	fmt.Println("After: ", reversedX.Pix[:16])
	reversedY := process.ReverseImageBits(imgY)
	if err := imageio.WriteGrayBMP(outputYReversed, headerX, reversedY.Pix, widthX, heightX); err != nil {
		panic(err)
	}
}
