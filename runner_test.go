package main

import (
	"github.com/stretchr/testify/mock"
	. "gopkg.in/check.v1"
)

type MockConverter struct {
	mock.Mock
}

func (_m *MockConverter) Convert(_a0 string, _a1 string, _a2 Rect) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, Rect) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *MockConverter) Size(_a0 string) (Rect, error) {
	ret := _m.Called(_a0)

	var r0 Rect
	if rf, ok := ret.Get(0).(func(string) Rect); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(Rect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type RunnerSuite struct {
	conv Converter
}

var _ = Suite(&RunnerSuite{
	conv: &RunnerSuiteConverter{},
})

type RunnerSuiteConverter struct {
}

func (r *RunnerSuiteConverter) Convert(infile string, outfile string, sourceRect Rect) error {
	return nil
}

func (r *RunnerSuiteConverter) Size(infile string) (Rect, error) {
	return Rect{60, 60}, nil
}

func (s *RunnerSuite) Testshould_auto_complete_filename_given_missing_in_template(c *C) {
	contents := NewContentsFromFile("fixtures/contents-no-sizes-no-files.json")
	r := NewRunner(contents, s.conv, "sourcefile.png")

	for _, meta := range contents.Images {
		c.Check(meta.Filename, Equals, "")
	}
	r.run()

	for _, meta := range contents.Images {
		c.Check(meta.Filename, Not(Equals), "")
	}
	c.Check(contents.Images[0].Filename, Equals, "sourcefile-iphone-20@2x.png")
	c.Check(contents.Images[6].Filename, Equals, "sourcefile-ipad-30@1x.png")
}

func (s *RunnerSuite) Test_should_leave_file_as_is_given_in_template(c *C) {
	contents := NewContentsFromFile("fixtures/with-files.json")
	r := NewRunner(contents, s.conv, "sourcefile.png")

	r.run()

	for _, meta := range contents.Images {
		c.Check(meta.Filename, Not(Equals), "")
	}
	c.Check(contents.Images[0].Filename, Equals, "Icon-Small@2x.png")
	c.Check(contents.Images[1].Filename, Equals, "sourcefile-iphone-20@3x.png")
}

func (s *RunnerSuite) Test_should_use_outdir_given_nonempty(c *C) {
	contents := NewContentsFromFile("fixtures/single-file.json")
	mk := &MockConverter{}
	mk.On("Size", "sourcefile.png").Return(Rect{42, 42}, nil)
	// 28 because it's 2/3
	mk.On("Convert", "sourcefile.png", "mydir/Icon-Small@2x.png", Rect{42, 42}).Return(nil)

	r := NewRunner(contents, mk, "sourcefile.png")
	r.OutDir = "mydir"

	r.run()

	mk.AssertExpectations(c)
}

func (s *RunnerSuite) Test_should_fill_in_computed_size_given_empty_size(c *C) {
	contents := NewContentsFromFile("fixtures/contents-no-sizes.json")
	r := NewRunner(contents, s.conv, "sourcefile.png")

	r.run()

	c.Check(contents.Images[0].Size, Equals, "20x20")
}

func (s *RunnerSuite) Test_should_convert_given_list_of_images(c *C) {
	contents := NewContentsFromFile("fixtures/contents-appicon.json")
	mk := &MockConverter{}
	mk.On("Size", "sourcefile.png").Return(Rect{42, 42}, nil)
	// 28 because it's 2/3
	mk.On("Convert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Times(12)

	r := NewRunner(contents, mk, "sourcefile.png")
	r.OutDir = "mydir"

	r.run()

	mk.AssertExpectations(c)
}
