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

	for _, s := range parser.Structs {
		fmt.Printf("%v\n", s.Name)
		for _, f := range s.Fields {
			fmt.Printf("%v %v\n", f.Name, f.Type)
		}
		fmt.Printf("\n")
	}
}
