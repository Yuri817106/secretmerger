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
	outputZ := "output/Z_combined.bmp"
	outputReversed := "output/Z_reversed.bmp"
	outputDoubleReversed := "output/Z_double_reversed.bmp"

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

	// 合併 X 和 Y 成 Z
	imgZ := process.Combine(imgX, imgY)

	// 對 Z 進行一次反轉
	reversedZ := process.ReverseImageBits(imgZ)

	// 再次反轉 Z
	doubleReversedZ := process.ReverseImageBits(reversedZ)

	// 保存 Z 、反轉後的 Z 和兩次反轉的 Z
	if err := imageio.WriteGrayBMP(outputZ, headerX, imgZ.Pix, widthX, heightX); err != nil {
		panic(err)
	}
	if err := imageio.WriteGrayBMP(outputReversed, headerX, reversedZ.Pix, widthX, heightX); err != nil {
		panic(err)
	}
	if err := imageio.WriteGrayBMP(outputDoubleReversed, headerX, doubleReversedZ.Pix, widthX, heightX); err != nil {
		panic(err)
	}

	fmt.Println("Saved combined image to", outputZ)
	fmt.Println("Saved reversed image to", outputReversed)
	fmt.Println("Saved double-reversed image to", outputDoubleReversed)
	fmt.Println("All images saved successfully.")

	process.SetReferenceImage(imgX)
	fmt.Println("PSNR (ReferenceImage X): ")
	fmt.Printf("Z : %.3f dB\n", process.PSNR(imgZ))

	process.SetReferenceImage(imgY)
	fmt.Println("PSNR (ReferenceImage Y): ")
	fmt.Printf("Reversed Z : %.3f dB\n", process.PSNR(reversedZ))
	fmt.Printf("Double Reversed Z : %.3f dB\n", process.PSNR(doubleReversedZ))
}
