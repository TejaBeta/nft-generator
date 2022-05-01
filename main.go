package main

import (
	"encoding/json"
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Metadata struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Hash       string     `json:"hash"`
	Date       time.Time  `json:"date"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Id    int    `json:"id"`
	Layer string `json:"layer"`
	Name  string `json:"name"`
}

func main() {
	var n int
	var l, f string
	flag.IntVar(&n, "n", 0, "Number to start from")
	flag.StringVar(&l, "layers", "./layers", "Layers location")
	flag.StringVar(&f, "final", "./final", "Final location")

	flag.Parse()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	multiLayers, err := readLayers(l)
	if err != nil {
		log.Error(err)
		return
	}
	compose(multiLayers, n, f)
}

func compose(m [][]string, n int, final string) {
	rand.Seed(time.Now().UnixNano())
	metadata := []Metadata{}
	g, h := make([]string, len(m)), make([]string, len(m))
	for i := 0; i < n; i++ {
		properties := []Property{}
	loop:
		for k, v := range m {
			r := rand.Intn(len(v))
			h[k], g[k] = strconv.Itoa(r), v[r]
			property := Property{r, strings.Split(strings.Split(v[r], "/")[1], "_")[1], strings.Split(strings.Split(v[r], "/")[2], ".")[0]}
			properties = append(properties, property)
		}
		if isHashExists(metadata, strings.Join(h, "")) {
			properties = properties[:len(properties)-1]
			goto loop
		}
		meta := Metadata{i, strconv.Itoa(i+1) + ".PNG", strings.Join(h, ""), time.Now(), properties}
		metadata = append(metadata, meta)
		generator(g, final+"/"+strconv.Itoa(i+1)+".PNG")
	}

	file, _ := json.MarshalIndent(metadata, "", " ")

	_ = ioutil.WriteFile(final+"/metadata.json", file, 0644)
}

func isHashExists(metadata []Metadata, hash string) bool {
	if len(metadata) > 0 {
		for _, v := range metadata {
			if hash == v.Hash {
				return true
			}
		}
	}
	return false
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

	result, err := os.Create(output)
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
