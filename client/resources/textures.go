package resources

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

// Only JPEGs supported for now
func ReadTexture(path ResPath) (*image.RGBA, error) {
	filePath := GameFilePath(path)
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	cardImage, err := jpeg.Decode(file)

	if err != nil {
		return nil, err
	}

	cardRGBA := image.NewRGBA(cardImage.Bounds())
	draw.Draw(cardRGBA, cardRGBA.Bounds(), cardImage, cardImage.Bounds().Min, draw.Src)

	return cardRGBA, nil
}
