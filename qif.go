package main

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Qif struct {
	Tags []Tag `json:"tags"`
}

func (q *Qif) AddTag(tag Tag) {
	q.Tags = append(q.Tags, tag)
}

func NewQif(tags []Tag) *Qif {
	return &Qif{Tags: tags}
}
