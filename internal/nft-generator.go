/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
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

var metaData []Metadata

func NFTGenerator(n int, l string, f string) {
	err := finalDir(f)
	if err != nil {
		log.Error(err)
		return
	}

	m, err := readLayers(l)
	if err != nil {
		log.Error(err)
		return
	}

	for i := 0; i <= n; i++ {
		h := compose(m)
		g := make([]string, len(m))
		for k, v := range h {
			g[k] = m[k][v]
		}
		generator(g, f+"/"+strconv.Itoa(i)+".PNG")
		meta := metaGenerator(m, h)
		meta.Id = i
		meta.Name = strconv.Itoa(i) + ".PNG"
		metaData = append(metaData, meta)
	}

	createMetaFile(metaData, f+"/metadata.json")
}

func metaGenerator(m [][]string, h []int) Metadata {
	properties := []Property{}
	for k, v := range h {
		properties = append(properties, Property{v, strings.Split(strings.Split(m[k][v], "/")[1], "_")[1], strings.Split(strings.Split(m[k][v], "/")[2], ".")[0]})
	}

	return Metadata{Hash: stringEncoder(h), Date: time.Now(), Properties: properties}
}

func stringEncoder(s []int) string {
	f := make([]string, len(s))
	for k, v := range s {
		f[k] = strconv.Itoa(v)
	}

	return fmt.Sprintf("%x", sha1.Sum([]byte(strings.Join(f, ""))))
}

func compose(m [][]string) []int {
	h := make([]int, len(m))
	rand.Seed(time.Now().UnixNano())

	for k, v := range m {
		h[k] = rand.Intn(len(v))
	}

	if isHashExists(metaData, stringEncoder(h)) {
		compose(m)
	}

	return h
}

func createMetaFile(metadata []Metadata, fileName string) error {
	file, err := json.MarshalIndent(metadata, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		return err
	}

	log.Info("Created metadata file at ", fileName)
	return nil
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
	log.Info(output)
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

func readLayers(dir string) ([][]string, error) {
	layers := []string{}
	multiLayers := [][]string{}

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.Name() != ".DS_Store" && f.Name() != strings.Split(dir, "/")[len(strings.Split(dir, "/"))-1] {
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

	for _, v := range multiLayers {
		for _, l := range v {
			err := validateLayer(strings.Split(l, "/")[1])
			if err != nil {
				return nil, err
			}
		}
	}

	return multiLayers, err
}

func finalDir(f string) error {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(f, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateLayer(l string) error {
	if len(strings.Split(l, "_")) < 2 || !strings.HasPrefix(l, "layer") {
		return errors.New("Layer is not named as per specifications: " + l)
	}

	return nil
}
