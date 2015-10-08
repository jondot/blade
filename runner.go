package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"path"
	"strings"
)

type Runner struct {
	converter        Converter
	contents         *Contents
	SourceFile       string
	OutDir           string
	GenerateContents bool
}

func NewRunner(contents *Contents, converter Converter, sourceFile string) *Runner {
	return &Runner{contents: contents, converter: converter, SourceFile: sourceFile}
}

func (f *Runner) run() {
	contents := f.contents

	sourceRect, err := f.converter.Size(f.SourceFile)
	if err != nil {
		log.Fatalf("Runner(File): cannot decode source file '%s' (%s).", f.SourceFile, err)
	}

	dim := NewDimensions()
	for i, _ := range contents.Images {
		meta := &contents.Images[i]
		rect := dim.Compute(contents, meta, sourceRect)

		// sync meta structure with the new data we've generated
		if meta.Size == "" {
			meta.Size = fmt.Sprintf("%vx%v", float64(rect.Width)/float64(meta.GetScale()), float64(rect.Height)/float64(meta.GetScale()))
		}
		if meta.Filename == "" {
			baseName := strings.TrimSuffix(path.Base(f.SourceFile), path.Ext(f.SourceFile))
			meta.Filename = meta.BuildFilename(baseName, rect)
		}

		outpath := path.Join(f.OutDir, meta.Filename)
		err = f.converter.Convert(f.SourceFile, outpath, rect)
		if err != nil {
			log.Printf("ERROR: cannot convert %s (%s)", outpath, err)
		}

		log.Infof("[%s] -> %v", f.OutDir, contents.Images[i])
	}
	log.Infof("[%s] %d images generated.", f.OutDir, len(contents.Images))

	if f.GenerateContents {
		err := contents.WriteToFile(path.Join(f.OutDir, "Contents.json"))
		if err != nil {
			log.Fatalf("Could not write Contents.json to %s (%s).", f.OutDir, err)
		}
		log.Infof("[%s] Wrote Contents.json.", f.OutDir)
	}
}
