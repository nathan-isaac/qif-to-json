package main

import (
	"bufio"
	"io"
)

type Parser struct {
	qif *Qif
}

func CreateParser() *Parser {
	return &Parser{
		qif: NewQif([]Tag{}),
	}
}

func (p *Parser) parse(handle io.Reader) {
	chunker := NewChunker()
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		line := scanner.Text()

		chunker.addLine(line)
	}

	for _, chunk := range chunker.GetChunks() {
		tp := TagParser{}
		tp.AddChunkToQif(chunk, p.qif)
	}
}
