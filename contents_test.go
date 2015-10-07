package main

import (
	. "gopkg.in/check.v1"
)

type ContentsSuite struct {
}

var _ = Suite(&ContentsSuite{})

func (cs *ContentsSuite) Test_should_build_filename(c *C) {
	ci := ContentsImage{}
	fname := ci.BuildFilename("base", Rect{42, 24})
	c.Check(fname, Equals, "base--42@1x.png")

	ci.Idiom = "iphone"
	fname = ci.BuildFilename("base", Rect{42, 24})
	c.Check(fname, Equals, "base-iphone-42@1x.png")

	ci.Scale = "1x"
	fname = ci.BuildFilename("base", Rect{42, 24})
	c.Check(fname, Equals, "base-iphone-42@1x.png")

	ci.Scale = "3x"
	fname = ci.BuildFilename("base", Rect{42, 24})
	c.Check(fname, Equals, "base-iphone-14@3x.png")
}

func (cs *ContentsSuite) Test_should_get_size(c *C) {
	ci := ContentsImage{}
	ci.Size = "42x24"
	w, h := ci.GetSize()
	c.Check(w, Equals, 42.0)
	c.Check(h, Equals, 24.0)

	ci.Size = "27.5x24.5"
	w, h = ci.GetSize()
	c.Check(w, Equals, 27.5)
	c.Check(h, Equals, 24.5)

	ci.Size = "27x24.5"
	w, h = ci.GetSize()
	c.Check(w, Equals, 27.0)
	c.Check(h, Equals, 24.5)

	ci.Size = "27.5x24"
	w, h = ci.GetSize()
	c.Check(w, Equals, 27.5)
	c.Check(h, Equals, 24.0)
}

func (cs *ContentsSuite) Test_should_get_scale(c *C) {
	ci := ContentsImage{}
	ci.Scale = "3x"
	s := ci.GetScale()
	c.Check(s, Equals, 3)

	ci.Scale = ""
	s = ci.GetScale()
	c.Check(s, Equals, 1)
}
