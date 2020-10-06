package main

type Line string

func (l Line) firstCharacterIs(s string) bool {
	return string(l)[0:1] == s
}

func (l Line) Value() string {
	return string(l)[1:]
}

type Chunk struct {
	Type string
	lines []Line
}

func (c *Chunk) addLine(line string) {
	c.lines = append(c.lines, Line(line))
}

type Chunker struct {
	lastChunk Chunk
	chunks    []Chunk
}

func (c *Chunker) addLine(line string) {
	if len(line) == 0 {
		return
	}

	c.lastChunk.addLine(line)

	firstChar := line[0:1]

	if firstChar == "^" {
		c.chunks = append(c.chunks, c.lastChunk)
		c.lastChunk = Chunk{}
	}
}

func (c *Chunker) GetChunks() []Chunk {
	return c.chunks
}

func NewChunker() *Chunker {
	return &Chunker{
		lastChunk: Chunk{},
		chunks: []Chunk{},
	}
}