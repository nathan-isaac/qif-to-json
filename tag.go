package main

type TagParser struct {}

func (tp *TagParser) AddChunkToQif(chunk Chunk, qif *Qif) {
	t := Tag{}

	for _, line := range chunk.lines {
		if line.firstCharacterIs("N") {
			t.Name = line.Value()
		}

		if line.firstCharacterIs("D") {
			t.Description = line.Value()
		}
	}

	qif.AddTag(t)
}
