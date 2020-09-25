package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		handler := strings.NewReader("")
		
		assert.Equal(t, qif{}, parse(handler))
	})

	t.Run("tags name", func(t *testing.T) {
		handler := strings.NewReader(`
!Type:Tag
NExample Name
^
`)

		assert.Equal(t, qif{
			[]tag{{
				name: "Example Name",
			}},
		}, parse(handler))
	})

	t.Run("tags description", func(t *testing.T) {
		handler := strings.NewReader(`
!Type:Tag
NExample Name
DExample Description
^
`)

		assert.Equal(t, qif{
			[]tag{{
				name: "Example Name",
				description: "Example Description",
			}},
		}, parse(handler))
	})

}
