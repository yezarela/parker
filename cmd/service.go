package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yezarela/parker/park"
)

// Commander represents commander
type Commander struct {
	Park *park.Park
}

// NewCommander will create a Commander instance
func NewCommander(Park *park.Park) *Commander {
	return &Commander{Park}
}

// ReadFromStdIn reads commands from stdin
func (c *Commander) ReadFromStdIn() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		c.ExecCommand(command)
	}
}

// ExecCommand process command input
func (c *Commander) ExecCommand(command string) {

	cmds := strings.Fields(
		strings.ToLower(strings.TrimSuffix(command, "\n")),
	)

	if len(cmds) < 1 {
		return
	}

	cmd := cmds[0]
	args := cmds[1:]
	output := ""

	switch cmd {
	case Init:
		output = c.doInit(args)

	case Park:
		output = c.doPark(args)

	case Leave:
		output = c.doLeave(args)

	case Exit:
		os.Exit(0)

	default:
		output = ErrUnknownCommand
	}

	if len(output) > 1 {
		fmt.Println(output)
	}
}

func (c *Commander) doInit(args []string) string {

	if slots, err := strconv.Atoi(args[0]); err != nil {
		return err.Error()
	} else {
		c.Park.InitSlots(slots)
		return fmt.Sprintf("Current slots: %d", slots)
	}
}

func (c *Commander) doPark(args []string) string {

	car := park.Car{
		RegNo: args[0],
		Color: args[1],
	}

	if id, err := c.Park.In(car); err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf("Slot %d is used", *id)
	}
}

func (c *Commander) doLeave(args []string) string {

	if id, err := strconv.Atoi(args[0]); err != nil {
		return err.Error()
	} else {
		if _, err := c.Park.Out(id); err != nil {
			return err.Error()
		} else {
			return fmt.Sprintf("Slot %d is free", id)
		}
	}
}