package main

import (
	"os"
)

func main() {
	analyze("data/quicken.QIF")
}

func analyze(file string) error {
	handle, err := os.Open(file)

	if err != nil {
		return err
	}
	defer handle.Close()
	p := Parser{}

	p.parse(handle)
	return nil
}
