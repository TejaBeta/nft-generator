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
	"errors"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestStringEncoder_A001(t *testing.T) {
	s := stringEncoder([]int{1, 2, 3, 4, 5})

	if s != "8cb2237d0679ca88db6464eac60da96345513964" {
		t.Errorf("Error while string encoding")
	}
}

func TestValidateLayer_A001(t *testing.T) {
	err := validateLayer("layera_something")
	if err != nil {
		log.Error(err)
		t.Errorf("Error while validating layer")
	}
}

func TestValidateLayer_A002(t *testing.T) {
	err := validateLayer("a_something")
	if err == nil {
		log.Error(err)
		t.Errorf("Error while validating layer")
	}
}

func TestValidateLayer_A003(t *testing.T) {
	err := validateLayer("layersomething")
	if err == nil {
		log.Error(err)
		t.Errorf("Error while validating layer")
	}
}

func TestFinalDir_A001(t *testing.T) {
	err := finalDir("../testdata/final")
	if err != nil {
		t.Errorf("Error creating final direcotry")
	} else {
		os.RemoveAll("../testdata/final")
	}
}

func TestFinalDir_A002(t *testing.T) {
	err := finalDir("./testdata/final")
	if err == nil {
		log.Error(err)
		t.Errorf("Error at finalDir function")
	}
}

func TestReadLayers_A001(t *testing.T) {
	m, err := readLayers("../testdata/properlayers/")
	if err != nil {
		log.Error(err)
		t.Errorf("Error while reading layers")
	}

	if len(m) != 2 {
		t.Errorf("Error while reading layers")
	}

	if len(m[0]) != 1 {
		t.Errorf("Error while reading layers")
	}

	if len(m[1]) != 1 {
		t.Errorf("Error while reading layers")
	}
}

func TestCompose_A001(t *testing.T) {
	m, _ := readLayers("../testdata/properlayers/")
	h, err := compose(m)
	if err != nil {
		t.Errorf("Error at compose function")
	}

	f := []int{0, 0}

	if len(h) != len(f) {
		t.Errorf("Error at compose function")
	}

	if h[0] != f[0] {
		t.Errorf("Error at compose function")
	}

	if h[1] != f[1] {
		t.Errorf("Error at compose function")
	}
}

func TestGenerator_A001(t *testing.T) {
	m, _ := readLayers("../testdata/properlayers/")
	h, err := compose(m)
	if err != nil {
		t.Errorf("Error at compose function")
	}
	g := make([]string, len(m))
	for k, v := range h {
		g[k] = m[k][v]
	}

	err = finalDir("../testdata/final")
	if err != nil {
		log.Error(err)
		t.Errorf("Error at creating final directory")
	}

	err = generator(g, "../testdata/final/test.png")
	if err != nil {
		t.Errorf("Error at Generator Function")
	}

	if _, err := os.Stat("../testdata/final/test.png"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Error at Generator Function unable to create a final file")
	} else {
		os.RemoveAll("../testdata/final")
	}
}

func TestMetaGenerator_A001(t *testing.T) {
	m, _ := readLayers("../testdata/properlayers/")
	h, err := compose(m)
	if err != nil {
		t.Errorf("Error at compose function")
	}

	meta := metaGenerator(m, h)

	if meta.Hash != "fb96549631c835eb239cd614cc6b5cb7d295121a" {
		t.Errorf("Error at hash while creating metadata")
	}

	if len(meta.Properties) != 2 {
		t.Errorf("Error at properties while creating metadata")
	}
}

func TestCreateMetaFile_A001(t *testing.T) {
	m, _ := readLayers("../testdata/properlayers/")
	h, err := compose(m)
	if err != nil {
		t.Errorf("Error at compose function")
	}

	err = finalDir("../testdata/final")
	if err != nil {
		log.Error(err)
		t.Errorf("Error at creating final directory")
	}

	createMetaFile([]Metadata{metaGenerator(m, h)}, "../testdata/final/metadata.json")

	if _, err := os.Stat("../testdata/final/metadata.json"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Error while creating metadata file")
	} else {
		os.RemoveAll("../testdata/final")
	}
}
