package main

import (
	"math"
)

type Dimensions struct {
}

type Rect struct {
	Width  float64
	Height float64
}

func NewDimensions() *Dimensions {
	return &Dimensions{}
}

func (d *Dimensions) Compute(contents *Contents, meta *ContentsImage, sourceRect Rect) Rect {
	if meta.Size != "" {
		return dimensionFromSize(meta)
	}

	//
	// If we don't get an explicit size we have to compute it.
	// 1. Find the biggest scale factor in the requested idiom group
	// 2. Assuming the source image is the largest possible image we have,
	//    compute a reduction scale factor (current / highest), and scale down the source.
	//    e.g if highest scale factor for iPad is x3, and the current image request is x2,
	//    we need to scale down the source image by 2/3.
	//

	// an idiom is just meta.Idiom, however with Apple Watch we have the same idiom with different
	// screen sizes. Feels like Apple patched this abstraction out with 'screenWidth'. So let's
	// compose a new idiom built from the original idiom and this screenWidth concept.
	idiom := meta.Idiom + meta.ScreenWidth
	highestFactor := 1.0

	for _, m := range contents.Images {
		if m.Idiom+m.ScreenWidth == idiom {
			highestFactor = math.Max(float64(highestFactor), float64(m.GetScale()))
		}
	}

	scaleDownFactor := float64(meta.GetScale()) / highestFactor
	//log.Printf("%v scale factor: %1.2f computed from %d and highest %2.0f", meta, scaleDownFactor, meta.GetScale(), highestFactor)
	return Rect{
		float64(sourceRect.Width) * scaleDownFactor,
		float64(sourceRect.Height) * scaleDownFactor,
	}

}

func dimensionFromSize(c *ContentsImage) Rect {
	w, h := c.GetSize()
	factor := float64(c.GetScale())

	return Rect{Width: factor * w, Height: factor * h}
}
