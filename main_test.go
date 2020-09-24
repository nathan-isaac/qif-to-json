package main

import (
"strings"
"testing"
)

func Test_analyze(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if err := doSomething(strings.NewReader("This is a test string")); (err != nil) != false {
			t.Errorf("analyze() error = %v", err)
		}
	})
}
