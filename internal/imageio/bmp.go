package imageio

import (
	"fmt"
	"os"
)

const bmpHeaderSize = 54 // BMP header size

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

func WriteGrayBMP(path string, header, pixels []byte) error {
	out := append(header, pixels...)
	return os.WriteFile(path, out, 0644)
}
