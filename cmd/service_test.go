package cmd

import (
	"testing"

	"github.com/yezarela/parker/park"
	"github.com/yezarela/parker/utils"
)

func TestExecCommand(t *testing.T) {
	p := park.NewPark()
	c := NewCommander(p)

	output := utils.CaptureOutput(func() {
		c.ExecCommand(Init + " 6")
	})

	expected := "Current slots: 6\n"
	if expected != output {
		t.Errorf("expect output to be '%s', got:  '%s'", expected, output)
	}
}

func TestExecCommandWithError(t *testing.T) {
	p := park.NewPark()
	c := NewCommander(p)

	output := utils.CaptureOutput(func() {
		c.ExecCommand("Die")
	})

	expected := "Unknown Command\n"
	if expected != output {
		t.Errorf("expect output to be '%s', got:  '%s'", expected, output)
	}
}
