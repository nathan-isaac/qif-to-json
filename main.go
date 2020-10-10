package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "", "File path to .qif file.")
	flag.Parse()

	handler, err := os.Open(*filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer handler.Close()

	parser := CreateParser()
	parser.Parse(handler)

	fmt.Println(parser.Json())
}
