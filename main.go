package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	layers, err := readLayers("./layers")
	if err != nil {
		log.Error(err)
		return
	}
	generator(layers)
}

func readLayers(dir string) ([]string, error) {
	layers := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			layers = append(layers, path)
		}
		return nil
	})
	return layers, err
}

func generator(images []string) {
	var decodedImages = make([]image.Image, len(images))

	for i, img := range images {
		decodedImages[i] = openAndDecode(img)
	}

	bounds := decodedImages[0].Bounds()
	newImage := image.NewRGBA(bounds)

	for _, img := range decodedImages {
		draw.Draw(newImage, img.Bounds(), img, image.ZP, draw.Over)
	}

	result, err := os.Create("./final/result.jpg")
	if err != nil {
		log.Fatalf("Failed to create: %s", err)
	}

	jpeg.Encode(result, newImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer result.Close()
	log.Info("File created at ./final/result.jpg")
}

func openAndDecode(imgPath string) image.Image {
	img, err := os.Open(imgPath)
	if err != nil {
		log.Fatalf("Failed to open %s", err)
	}

	decoded, err := png.Decode(img)
	if err != nil {
		log.Fatalf("Failed to decode %s", err)
	}
	defer img.Close()

	return decoded
}
