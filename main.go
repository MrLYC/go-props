package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var err error

	log.SetOutput(ioutil.Discard)
	err = ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	if Config.WithLog {
		log.SetOutput(os.Stderr)
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
