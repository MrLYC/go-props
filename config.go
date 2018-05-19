package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
)

// ConfigType : global config type
type ConfigType struct {
	Files []string
}

// Config : global config
var Config ConfigType

func toAvaliableGoFile(f string) (string, error) {
	ext := filepath.Ext(f)
	if ext != ".go" {
		return "", ErrNotGoFile
	}
	file, err := filepath.Abs(f)
	return file, err
}

// ParseConfig :
func ParseConfig() error {
	var directory, file string
	flag.StringVar(&directory, "p", "", "package path")
	flag.StringVar(&file, "f", "", "file path")
	flag.Parse()

	if file != "" {
		file, err := toAvaliableGoFile(file)
		if err != nil {
			return err
		}
		Config.Files = append(Config.Files, file)
	}

	if directory != "" {
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			file, err = toAvaliableGoFile(filepath.Join(directory, f.Name()))
			if err == ErrNotGoFile {
				continue
			} else if err != nil {
				return err
			}
			Config.Files = append(Config.Files, file)
		}
	}

	return nil
}
