package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var err error

	log.SetOutput(os.Stderr)
	err = ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	if Config.LogDisabled {
		log.SetOutput(ioutil.Discard)
	}

	parser := NewParser()
	for _, file := range Config.Files {
		log.Printf("parsing file: %v", file)
		err = parser.ParseFile(file)
		if err != nil {
			log.Panic(err)
		}
	}

	log.Printf("preparing")
	gen := parser.Prepare()
	fmt.Println(gen.Generate())
}
