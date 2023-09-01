### Synopsis

* Go library to capture desktop screen.
* Supported GOOS: windows.
* `cgo` free except

It captures a screenshot of the screen or a specific rectangular area of the screen using the Windows API. The CaptureScreen function captures the entire screen, while the CaptureRect function captures a specific rectangular area. The code uses various Windows API functions such as GetDC, CreateCompatibleDC, CreateDIBSection, and BitBlt to capture the screenshot and store it in an image.RGBA value. You're able to capture a screenshot of the whole Desktop with this program.

# Example of usage:

```go
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

```

### End

If you encounter any issues or have suggestions, please open an issue on GitHub or contact me on Twitter at @n1k7l4i. Your feedback helps improve this project.
