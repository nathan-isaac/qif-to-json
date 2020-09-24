package main

import (
"bufio"
"io"
"os"
)

func main() {
	analyze("data/quicken.QIF")
}

func analyze(file string) error {
	handle, err := os.Open(file)

	if err != nil {
		return err
	}
	defer handle.Close()
	return doSomething(handle)
}

func doSomething(handle io.Reader) error {
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		println(scanner.Text())
	}
	return nil
}