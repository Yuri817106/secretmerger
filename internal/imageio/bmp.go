package imageio

import (
	"fmt"
	"os"
	"encoding/binary"
)

const (
    bmpHeaderSize = 54
    paletteSize   = 256 * 4
)

func ReadGrayBMP(path string) ([]byte, []byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	if len(file) < bmpHeaderSize {
		return nil, nil, fmt.Errorf("not a valid BMP file")
	}

	header := file[:bmpHeaderSize]
	pixels := file[bmpHeaderSize:]
	return header, pixels, nil
}

func FixBMPHeader(header []byte, width, height int) {
    // 設置調色板偏移量
    binary.LittleEndian.PutUint32(header[10:], bmpHeaderSize+paletteSize)
    
    // 更新位深度為8-bit
    binary.LittleEndian.PutUint16(header[28:], 8)
    
    // 設置調色板顏色數
    binary.LittleEndian.PutUint32(header[46:], 256)
}

func WriteGrayBMP(path string, header, pixels []byte, width int, height int) error {
    // 添加調色板
    palette := make([]byte, paletteSize)
    for i := 0; i < 256; i++ {
        palette[i*4] = byte(i)     // B
        palette[i*4+1] = byte(i)   // G
        palette[i*4+2] = byte(i)   // R
        // i*4+3 保持0
    }
    
    // 計算行填充
    bytesPerRow := width
    padding := (4 - (bytesPerRow % 4)) % 4
    paddedPixels := make([]byte, 0, len(pixels) + (padding * height))
    
    // 重新排列像素並添加填充
    for y := 0; y < height; y++ {
        start := y * width
        end := start + width
        paddedPixels = append(paddedPixels, pixels[start:end]...)
        paddedPixels = append(paddedPixels, make([]byte, padding)...)
    }
    
    // 組合最終文件
    fullHeader := append(header[:bmpHeaderSize], palette...)
    return os.WriteFile(path, append(fullHeader, paddedPixels...), 0644)
}
