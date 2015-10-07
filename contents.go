package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
)

type Contents struct {
	Images []ContentsImage `json:"images"`
	Info   struct {
		Author  string `json:"author"`
		Version int    `json:"version"`
	} `json:"info"`
}

type ContentsImage struct {
	Filename    string `json:"filename"`
	Idiom       string `json:"idiom"`
	Scale       string `json:"scale"`
	Size        string `json:"size"`
	ScreenWidth string `json:"screenWidth,omitempty"`
}

func (ci *ContentsImage) GetScale() int {
	if ci.Scale == "" {
		return 1
	}

	factor, err := strconv.Atoi(ci.Scale[0:1])
	if err != nil {
		log.Fatalf("Converter(Resize): cannot parse scale %s (%s).", ci.Size, err)
	}
	return factor
}

func (ci *ContentsImage) GetSize() (float64, float64) {
	a := strings.Split(ci.Size, "x")
	w, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		log.Fatalf("Converter(Resize): cannot parse width %s (%s).", a[0], err)
	}
	h, err := strconv.ParseFloat(a[1], 64)
	if err != nil {
		log.Fatalf("Converter(Resize): cannot parse height %s (%s).", a[0], err)
	}

	return w, h
}

func (ci *ContentsImage) BuildFilename(base string, rect Rect) string {
	scale := ci.GetScale()
	return fmt.Sprintf("%s-%s-%d@%dx.png", base, ci.Idiom, int(float64(rect.Width)/float64(scale)), scale)
}

func NewContentsFromFile(path string) *Contents {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Contents: cannot read from %s (%s).", path, err)
	}
	return NewContentsFromString(data)
}

func NewContentsFromString(data []byte) *Contents {
	var contents Contents
	json.Unmarshal(data, &contents)
	return &contents
}

func (c *Contents) WriteToFile(file string) error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
