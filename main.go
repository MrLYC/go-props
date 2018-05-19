package main

import (
	"fmt"
)

func main() {
	var err error
	err = ParseConfig()
	if err != nil {
		panic(err)
	}

	parser := NewParser()
	for _, file := range Config.Files {
		err = parser.ParseFile(file)
		if err != nil {
			panic(err)
		}
	}

	gen := parser.Prepare()
	fmt.Println(gen.Generate())
}
