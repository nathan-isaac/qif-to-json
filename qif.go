package main

type Qif struct {
	tags []Tag
}

func (q *Qif) AddTag(tag Tag) {
	q.tags = append(q.tags, tag)
}

func NewQif(tags []Tag) *Qif {
	return &Qif{tags: tags}
}
