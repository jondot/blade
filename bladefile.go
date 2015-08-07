package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Bladefile struct {
	Blades []Blade `yaml:"blades"`
}

func (bf *Bladefile) Exists() bool {
	if _, err := os.Stat("Bladefile"); os.IsNotExist(err) {
		return false
	}
	return true
}

func (bf *Bladefile) Load() error {
	data, err := ioutil.ReadFile("Bladefile")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(data), bf)
	if err != nil {
		return err
	}
	return nil
}
