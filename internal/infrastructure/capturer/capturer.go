package capturer

import (
	"errors"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath
func save(img *image.RGBA, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// var imageQuality = 100
	// fileSizeEstimated := len(img.Pix) * 4 // Each pixel is 4 bytes (RGBA)
	// fileSizeLimit := 500 * 1024           // Target file size in bytes

	// if fileSizeEstimated > fileSizeLimit {
	// 	imageQuality = 100 * (fileSizeLimit / fileSizeEstimated)
	// }

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}

func GetMergedScreenFilePath() string {
	fileName := "merged-screen.jpeg"

	return filepath.Join(os.TempDir(), fileName)
}

func CaptureMergedScreen() error {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		return errors.New("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		/** Capture indivisual screens
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return err
		}
		fileName := fmt.Sprintf("%d_%dx%d.jpeg", i, bounds.Dx(), bounds.Dy())
		save(img, fileName)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		*/
	}

	// Capture all desktop region into an image.
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		return err
	}

	destinationPath := GetMergedScreenFilePath()

	return save(img, destinationPath)
}
