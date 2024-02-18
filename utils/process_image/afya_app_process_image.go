package process_image

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/exec"

	"github.com/EventHubzTz/event_hub_service/app/models"
)

var ProcessImage = newProcessImage()

type processImage struct {
}

func newProcessImage() processImage {
	return processImage{}
}

func (_ processImage) GetImageDimension(imagePath string) (*models.ImageDimension, error) {
	var dimension models.ImageDimension
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(0, 0)
	image, _, err := image.DecodeConfig(file) // Image Struct
	if err != nil {
		return nil, err
	}
	dimension.Width = image.Width
	dimension.Height = image.Height
	return &dimension, nil
}

func (pI processImage) GetImageAspectRatios(imagePath string) (string, error) {
	dimension, err := pI.GetImageDimension(imagePath)
	if err != nil {
		return "1.5", err
	}
	x := fmt.Sprintf("%.2f", (float32(dimension.Width) / float32(dimension.Height)))
	return x, nil
}

func (_ processImage) DeleteImage(imagePath string) error {
	err := os.Remove(imagePath)
	if err != nil {
		return err

	}
	return nil
}

func GenerateThumbnail(videoPath string, thumbnailPath string, time string) error {
	cmd := exec.Command("ffmpeg",
		"-i", videoPath,
		"-ss", time, // Time in the format "hh:mm:ss" or seconds
		"-vframes", "1",
		thumbnailPath,
	)
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
