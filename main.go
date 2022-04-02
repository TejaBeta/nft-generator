package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.Println("Hello")
	generator()
}

func generator() {
	var images = []string{"bg_0", "edge_0", "edge_1", "edge_2", "a_0", "b_0", "c_0", "d_0", "e_0", "f_0", "g_0", "h_0", "i_0"}
	var decodedImages = make([]image.Image, len(images))

	for i, img := range images {
		decodedImages[i] = openAndDecode("./layers/" + img + ".PNG")
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

	jpeg.Encode(result, newImage, &jpeg.Options{jpeg.DefaultQuality})
	defer result.Close()
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
