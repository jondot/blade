package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

var interpolations = map[string]resize.InterpolationFunction{
	"l3": resize.Lanczos3,
	"l2": resize.Lanczos2,
	"n":  resize.NearestNeighbor,
	"bc": resize.Bicubic,
	"bl": resize.Bilinear,
	"mn": resize.MitchellNetravali,
}

type ResizeConverter struct {
	Interpolation string
}

func NewResizeConverter() *ResizeConverter {
	return &ResizeConverter{}
}

func (r *ResizeConverter) Size(inFile string) (Rect, error) {
	in, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Resize: cannot open file '%s' (%s).", inFile, err)
	}

	source, _, err := image.Decode(in)
	if err != nil {
		return Rect{}, err
	}

	return Rect{
		float64(source.Bounds().Dx()),
		float64(source.Bounds().Dy()),
	}, nil
}

func (r *ResizeConverter) Convert(inFile string, outFile string, rect Rect) error {
	in, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("Runner(File): cannot open source file '%s' for reading (%s).", inFile, err)
	}
	defer in.Close()

	out, err := os.Create(outFile)
	if err != nil {
		log.Printf("ERROR: cannot create new file %s (%s)", outFile, err)
	}
	defer out.Close()

	source, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	interp, ok := interpolations[r.Interpolation]
	if !ok {
		interp = interpolations["l3"]
	}

	resized := resize.Resize(uint(rect.Width), uint(rect.Height), source, interp)
	png.Encode(out, resized)
	return nil
}
