package process

import (
    "image"
    "image/color"
)

// Combine 將兩張灰階影像 X 和 Y 合併為 Z
func Combine(x, y *image.Gray) *image.Gray {
    if x.Bounds() != y.Bounds() {
        panic("images must have the same dimensions")
    }

    bounds := x.Bounds()
    z := image.NewGray(bounds)

    for yCoord := bounds.Min.Y; yCoord < bounds.Max.Y; yCoord++ {
        for xCoord := bounds.Min.X; xCoord < bounds.Max.X; xCoord++ {
            xPixel := x.GrayAt(xCoord, yCoord).Y
            yPixel := y.GrayAt(xCoord, yCoord).Y
            
            // 合併方式：X7X6X5X4Y4Y5Y6Y7
            zPixel := (xPixel & 0xF0) | (yPixel >> 4)
            z.SetGray(xCoord, yCoord, color.Gray{Y: zPixel})
        }
    }

    return z
}