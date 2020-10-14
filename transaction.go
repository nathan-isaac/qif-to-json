package main

import "time"

type TransactionMapper struct {}

func (tp *TransactionMapper) AddChunkToQif(chunk Chunk, qif *Qif) {
	t := Transaction{}

	for _, line := range chunk.lines {
		if line.firstCharacterIs("D") {
			// parse line for date
			t.Date = time.Date(2014, 4, 10, 0,0,0,0, time.Local)
		}
	}

	qif.AddTransaction(t)
}
