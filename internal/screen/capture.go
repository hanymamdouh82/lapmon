// screen/capture.go
package screen

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

func TakeScreenshot(outputDir string) error {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}

	now := time.Now().Format("2006-01-02_15-04-05")
	filename := filepath.Join(outputDir, "", fmt.Sprintf("screenshot_%s.png", now))

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
