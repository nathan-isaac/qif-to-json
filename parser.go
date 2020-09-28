package main

import (
	"bufio"
	"io"
)

type Tag struct {
	name string
	description string
}

type Parser struct {
	tags []Tag
}

func CreateParser() *Parser {
	return &Parser{tags: []Tag{}}
}

func (p *Parser) parse(handle io.Reader) {
	tag := Tag{}

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
			p.tags = append(p.tags, tag)
		}
	}
}
