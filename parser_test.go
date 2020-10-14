package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestParsingTags(t *testing.T) {
	var tagTests = []struct {
		name string
		in   string
		out  []Tag
	}{
		{
			"empty string",
			"",
			[]Tag{},
		},
		{
			"ignore types without a parser",
			`
!Type:NoParser
NExample Name
^
`,
			[]Tag{},
		},
		{
			"tag with Name",
			`
!Type:Tag
NExample Name
^
`,
			[]Tag{{
				Name: "Example Name",
			}},
		},
		{
			"tag with Description",
			`
!Type:Tag
DExample Description
^
`,
			[]Tag{{
				Description: "Example Description",
			}},
		},
		{
			"tag with Description",
			`
!Type:Tag
DExample Description
^
`,
			[]Tag{{
				Description: "Example Description",
			}},
		},
		{
			"tag with Name and Description",
			`
!Type:Tag
NExample Name
DExample Description
^
`,
			[]Tag{{
				Name:        "Example Name",
				Description: "Example Description",
			}},
		},
		{
			"multiple Tags",
			`
!Type:Tag
NTag 1
^
NTag 2
DTag 2 Description
^
`,
			[]Tag{{
				Name: "Tag 1",
			}, {
				Name:        "Tag 2",
				Description: "Tag 2 Description",
			}},
		},
	}

	for _, tt := range tagTests {
		t.Run(tt.name, func(t *testing.T) {
			handler := strings.NewReader(tt.in)
			parser := CreateParser()
			parser.Parse(handler)

			assert.Equal(t, tt.out, parser.qif.Tags)
		})
	}
}

func TestParsingTransactions(t *testing.T) {
	var transactionTests = []struct {
		name string
		in   string
		out  []Transaction
	}{
		{
			"empty string",
			"",
			[]Transaction{},
		},
		{
			"ignore types without a parser",
			`
!Type:NoParser
NExample Name
^
`,
			[]Transaction{},
		},
		{
			"transaction with date",
			`
!Type:Bank
D4/10'14
^
`,
			[]Transaction{{
				Date: time.Date(2014, 4, 10, 0, 0, 0, 0, time.Local),
			}},
		},
	}

	for _, tt := range transactionTests {
		t.Run(tt.name, func(t *testing.T) {
			handler := strings.NewReader(tt.in)
			parser := CreateParser()
			parser.Parse(handler)

			assert.Equal(t, tt.out, parser.qif.Transactions)
		})
	}
}
