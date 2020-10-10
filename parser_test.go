package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
