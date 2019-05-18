package utils

import (
	"bytes"
	"io"
	"os"
)

// CaptureOutput captures output from stdin
func CaptureOutput(f func()) string {

	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
