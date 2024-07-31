package capturer

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func GetMergedScreenFilePath() string {
	fileName := "merged-screen.png"

	return filepath.Join(os.TempDir(), fileName)
}

func CaptureMergedScreen() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		/** Capture indivisual screens
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		save(img, fileName)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		*/
	}

	// Capture all desktop region into an image.
	// fmt.Printf("%v\n", all)
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}

	destinationPath := GetMergedScreenFilePath()
	save(img, destinationPath)
}
