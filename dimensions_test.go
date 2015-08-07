package main

import (
	. "gopkg.in/check.v1"
)

type DimensionsSuite struct {
	dim *Dimensions
}

var _ = Suite(&DimensionsSuite{
	dim: NewDimensions(),
})

func (s *DimensionsSuite) Test_should_compute_size_based_on_existing_scales_given_no_size(c *C) {
	contents := NewContentsFromFile("fixtures/contents-no-sizes.json")
	// iphone
	r := s.dim.Compute(
		contents,
		&contents.Images[0],
		Rect{60, 60},
	)
	c.Check(r.Width, Equals, uint(40))
	c.Check(r.Height, Equals, uint(40))

	// ipad
	r = s.dim.Compute(
		contents,
		&contents.Images[2],
		Rect{60, 60},
	)
	c.Check(r.Width, Equals, uint(60))
	c.Check(r.Height, Equals, uint(60))

}

func (s *DimensionsSuite) Test_should_use_size_and_scale_when_given_size_explicitely(c *C) {
	contents := NewContentsFromFile("fixtures/contents-appicon.json")
	r := s.dim.Compute(
		contents,
		&contents.Images[0],
		Rect{42, 42},
	)
	c.Check(r.Width, Equals, uint(29*2))
	c.Check(r.Height, Equals, uint(29*2))
}
