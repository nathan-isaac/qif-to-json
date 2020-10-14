package main

import "time"

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Transaction struct {
	Date time.Time
}

type Qif struct {
	Tags         []Tag `json:"tags"`
	Transactions []Transaction
}

func (q *Qif) AddTag(tag Tag) {
	q.Tags = append(q.Tags, tag)
}

func (q *Qif) AddTransaction(t Transaction) {
	q.Transactions = append(q.Transactions, t)
}

func NewQif(tags []Tag) *Qif {
	return &Qif{
		Tags:         tags,
		Transactions: []Transaction{},
	}
}
