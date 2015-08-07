package main

import (
	log "github.com/Sirupsen/logrus"
)

type DryrunConverter struct {
}

func NewDryrunConverter() *DryrunConverter {
	return &DryrunConverter{}
}
func (d *DryrunConverter) Size(in string) (Rect, error) {
	return Rect{42, 42}, nil
}

func (d *DryrunConverter) Convert(in string, out string, toRect Rect) error {
	log.Infof("(dryrun) Converting %s to %s using %v\n", in, out, toRect)
	return nil
}
