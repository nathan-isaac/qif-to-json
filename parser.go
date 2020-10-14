package main

import (
	"bufio"
	"encoding/json"
	"io"
)

// Maybe call this TypeMapper?
type TypeParser interface {
	AddChunkToQif(chunk Chunk, qif *Qif)
}

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

	parsers := map[string]TypeParser{
		"Tag": &TagParser{},
		"Bank": &TransactionMapper{},
	}

	for _, chunk := range chunker.GetChunks() {
		if val, ok := parsers[chunk.Type]; ok {
			val.AddChunkToQif(chunk, p.qif)
		}
	}
}

func (p *Parser) Json() string {
	qif, _ := json.MarshalIndent(p.qif, "", "  ")
	return string(qif)
}
