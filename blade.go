package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"path"
)

type Blade struct {
	Template        string `yaml:"template"`
	Interpolation   string `yaml:"interpolation"`
	Source          string `yaml:"source"`
	Out             string `yaml:"out"`
	IncludeContents bool   `yaml:"contents"`
	DryRun          bool
	Mount           string `yaml:"mount"`
}

func (b *Blade) Run() {
	if b.Mount != "" {
		template := path.Join(b.Mount, "Contents.json")
		if _, err := os.Stat(template); os.IsNotExist(err) {
			log.Fatalf("A mount must point to an image catalog (Contents.json missing)")
		}
		b.Out = b.Mount
		b.Template = template
	}

	c := NewContentsFromFile(b.Template)

	var converter Converter

	if b.DryRun {
		converter = NewDryrunConverter()
	} else {
		cv := NewResizeConverter()
		cv.Interpolation = b.Interpolation
		converter = cv
	}

	r := NewRunner(c, converter, b.Source)
	if b.Out != "" {
		err := os.MkdirAll(b.Out, 0755)
		if err != nil {
			log.Fatalf("Cannot create output directory '%s' (%s)", b.Out, err)
		}
		r.OutDir = b.Out
	}

	if b.IncludeContents {
		r.GenerateContents = true
	}

	r.run()

}
