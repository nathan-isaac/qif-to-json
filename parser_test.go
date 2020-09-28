package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parsing_tag(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		handler := strings.NewReader("")
		parser := CreateParser()
		parser.parse(handler)

		assert.Equal(t, []Tag{}, parser.tags)
	})

	t.Run("tags name", func(t *testing.T) {
		handler := strings.NewReader(`
!Type:Tag
NExample Name
^
`)
		parser := CreateParser()
		parser.parse(handler)

		assert.Equal(t, []Tag{{
			name: "Example Name",
		}}, parser.tags)
	})

	t.Run("tags description", func(t *testing.T) {
		handler := strings.NewReader(`
!Type:Tag
DExample Description
^
`)
		parser := CreateParser()
		parser.parse(handler)

		assert.Equal(t, []Tag{{
			description: "Example Description",
		}}, parser.tags)
	})

	t.Run("tag with name and description", func(t *testing.T) {
		handler := strings.NewReader(`
!Type:Tag
NExample Name
DExample Description
^
`)
		parser := CreateParser()
		parser.parse(handler)

		assert.Equal(t, []Tag{{
			name: "Example Name",
			description: "Example Description",
		}}, parser.tags)
	})

	t.Run("multiple tags", func(t *testing.T) {
		handler := strings.NewReader(`
!Type:Tag
NTag 1
^
NTag 2
DTag 2 Description
^
`)
		parser := CreateParser()
		parser.parse(handler)

		assert.Equal(t, []Tag{{
			name: "Tag 1",
		}, {
			name: "Tag 2",
			description: "Tag 2 Description",
		}}, parser.tags)
	})

}
