package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ConfigType : global config type
type ConfigType struct {
	Files []string
	Stdin bool

	Package string
	LineSep string

	WithoutPrivateStruct bool
	WithoutPublicField   bool

	TagName string

	WithLog bool
}

// Config : global config
var Config ConfigType

func toAvaliableGoFile(f string, force bool) (string, error) {
	ext := filepath.Ext(f)
	if ext != ".go" && !force {
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

	var stdin bool
	flag.BoolVar(&stdin, "stdin", false, "from stdin")

	flag.StringVar(&(Config.Package), "with_package", "", "with package")
	flag.StringVar(&(Config.LineSep), "line_sep", "", "generate code line sep")

	flag.BoolVar(&(Config.WithoutPrivateStruct), "with_private_struct", false, "with private struct")
	flag.BoolVar(&(Config.WithoutPublicField), "with_public_field", false, "with public attributes")

	flag.StringVar(&(Config.TagName), "tag_name", "props", "struct property tag name")

	flag.BoolVar(&(Config.WithLog), "v", false, "with log")
	flag.BoolVar(&(Config.WithLog), "with_log", false, "with log")

	flag.Parse()

	if file != "" {
		file, err := toAvaliableGoFile(file, true)
		if err != nil {
			return err
		}
		Config.Files = append(Config.Files, file)
	}

	if stdin {
		tempFile, err := ioutil.TempFile("", "props")
		if err != nil {
			return err
		}
		_, err = io.Copy(tempFile, os.Stdin)
		if err != nil {
			return err
		}

		err = tempFile.Sync()
		if err != nil {
			return err
		}

		log.Printf("tempfile: %v", tempFile.Name())
		Config.Files = append(Config.Files, tempFile.Name())
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
			file, err = toAvaliableGoFile(filepath.Join(directory, f.Name()), false)
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
