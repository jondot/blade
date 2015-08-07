package main

type Converter interface {
	Convert(string, string, Rect) error
	Size(string) (Rect, error)
}
