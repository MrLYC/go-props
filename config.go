package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
)

// ConfigType : global config type
type ConfigType struct {
	Files []string

	LineSep string

	WithPrivateStruct bool
	WithPublicField   bool

	TagName string

	LogDisabled bool
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
	flag.StringVar(&(Config.LineSep), "line_sep", "\n", "generate code line sep")

	flag.BoolVar(&(Config.WithPrivateStruct), "with_private_struct", false, "with private struct")
	flag.BoolVar(&(Config.WithPublicField), "with_public_field", false, "with public attributes")

	flag.StringVar(&(Config.TagName), "tag_name", "props", "struct property tag name")

	flag.BoolVar(&(Config.LogDisabled), "log_disabled", false, "disable logger")

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
