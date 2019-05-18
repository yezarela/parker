package utils

import (
	"fmt"
	"testing"
)

func TestCaptureOutput(t *testing.T) {

	output := CaptureOutput(func() {
		fmt.Println("Ava cadavra")
	})

	expected := "Ava cadavra\n"
	if expected != output {
		t.Errorf("expect output to be '%s', got:  '%s'", expected, output)
	}
}
