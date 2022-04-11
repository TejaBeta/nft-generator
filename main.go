package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var n int
	var g bool
	flag.IntVar(&n, "n", 0, "Number to start from")
	flag.BoolVar(&g, "g", true, "Should generate or not")

	flag.Parse()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	multiLayers, err := readLayers("./layers")
	if err != nil {
		log.Error(err)
		return
	}
	composer(multiLayers, n*1024, g)
}

func composer(m [][]string, n int, g bool) {
	counter := 0
	imageCounter := 0
	start := time.Now()
	for _, l0 := range m[0] {
		for _, l1 := range m[1] {
			for _, l2 := range m[2] {
				for _, l3 := range m[3] {
					for _, la := range m[4] {
						for _, lb := range m[5] {
							for _, lc := range m[6] {
								for _, ld := range m[7] {
									for _, le := range m[8] {
										for _, lf := range m[9] {
											for _, lg := range m[10] {
												for _, lh := range m[11] {
													for _, li := range m[12] {
														for _, lj := range m[13] {
															for _, lk := range m[14] {
																_ = []string{l0, l1, l2, l3, la, lb, lc, ld, le, lf, lg, lh, li, lj, lk}
																if counter >= n && counter < (n+1025) {
																	if g {
																		generator([]string{l0, l1, l2, l3, la, lb, lc, ld, le, lf, lg, lh, li, lj, lk}, strconv.Itoa(counter)+".PNG")
																	}
																	imageCounter = imageCounter + 1
																}
																counter = counter + 1
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	duration := time.Since(start)
	log.Println("Images Created: ", imageCounter)
	log.Println("Total Duration: ", duration)
	log.Println("Total possibilities: ", counter)
}

func readLayers(dir string) ([][]string, error) {
	layers := []string{}
	multiLayers := [][]string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.Name() != ".DS_Store" && f.Name() != "layers" {
			if !f.IsDir() {
				layers = append(layers, path)
			} else {
				if len(layers) > 0 {
					multiLayers = append(multiLayers, layers)
					layers = nil
				}
			}
		}
		return nil
	})
	multiLayers = append(multiLayers, layers)
	return multiLayers, err
}

func generator(images []string, output string) {
	var decodedImages = make([]image.Image, len(images))

	for i, img := range images {
		decodedImages[i] = openAndDecode(img)
	}

	bounds := decodedImages[0].Bounds()
	newImage := image.NewRGBA(bounds)

	for _, img := range decodedImages {
		draw.Draw(newImage, img.Bounds(), img, image.ZP, draw.Over)
	}

	result, err := os.Create("./final/" + output)
	if err != nil {
		log.Fatalf("Failed to create: %s", err)
	}

	jpeg.Encode(result, newImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer result.Close()
	log.Printf("%s", output)
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
