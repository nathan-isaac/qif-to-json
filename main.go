package main

import (
	"bufio"
	"io"
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

	parse(handle)
	return nil
}

type tag struct {
	name string
	description string
}

type qif struct {
	tags []tag
}

func parse(handle io.Reader) qif {
	output := qif{}
	tag := tag{}

	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		first_char := line[0:1]

		if first_char == "N" {
			tag.name = line[1:]
		}
		if first_char == "D" {
			tag.description = line[1:]
		}

		if first_char == "^" {
			output.tags = append(output.tags, tag)
		}
	}

	return output
}

