package main

import (
	"bufio"
	"encoding/json"
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

func (p *Parser) Parse(handle io.Reader) {
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

func (p *Parser) Json() string {
	qif, _ := json.MarshalIndent(p.qif, "", "  ")
	return string(qif)
}
