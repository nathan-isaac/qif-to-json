package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsingTags(t *testing.T) {
	var tagTests = []struct {
		name  string
		in  string
		out []Tag
	}{
		{
			"empty string",
			"",
			[]Tag{},
		},
		{
			"tag with name",
			`
!Type:Tag
NExample Name
^
`,
		[]Tag{{
				name: "Example Name",
			}},
		},
		{
			"tag with description",
			`
!Type:Tag
DExample Description
^
`,
		[]Tag{{
				description: "Example Description",
			}},
		},
		{
			"tag with description",
			`
!Type:Tag
DExample Description
^
`,
			[]Tag{{
				description: "Example Description",
			}},
		},
		{
			"tag with name and description",
			`
!Type:Tag
NExample Name
DExample Description
^
`,
			[]Tag{{
				name: "Example Name",
				description: "Example Description",
			}},
		},
		{
			"multiple tags",
			`
!Type:Tag
NTag 1
^
NTag 2
DTag 2 Description
^
`,
			[]Tag{{
				name: "Tag 1",
			}, {
				name: "Tag 2",
				description: "Tag 2 Description",
			}},
		},
	}

	for _, tt := range tagTests {
		t.Run(tt.name, func(t *testing.T) {
			handler := strings.NewReader(tt.in)
			parser := CreateParser()
			parser.parse(handler)

			assert.Equal(t, tt.out, parser.tags)
		})
	}
}
