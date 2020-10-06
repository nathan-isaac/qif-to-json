package main

type Tag struct {
	name        string
	description string
}

type TagParser struct {}

func (tp *TagParser) AddChunkToQif(chunk Chunk, qif *Qif) {
	t := Tag{}

	for _, line := range chunk.lines {
		if line.firstCharacterIs("N") {
			t.name = line.Value()
		}

		if line.firstCharacterIs("D") {
			t.description = line.Value()
		}
	}

	qif.AddTag(t)
}
