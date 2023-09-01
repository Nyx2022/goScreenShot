package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	img, err := CaptureScreen()
	if err != nil {
		fmt.Println("Failed to capture screen, error: ", err)
	}

	filePath := "Screenshot.png"

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file, error: ", err)
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		fmt.Println("Failed to encode image, error: ", err)
		return
	}
}
