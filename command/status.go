package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/command/server/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

// StatusCommand is the command to output the status of the client
type StatusCommand struct {
	*Meta2
}

// Help implements the cli.Command interface
func (p *StatusCommand) Help() string {
	return `Usage: bor status

  Output the status of the client`
}

// Synopsis implements the cli.Command interface
func (c *StatusCommand) Synopsis() string {
	return "Output the status of the client"
}

// Run implements the cli.Command interface
func (c *StatusCommand) Run(args []string) int {
	flags := c.NewFlagSet("status")
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	borClt, err := c.BorConn()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	status, err := borClt.Status(context.Background(), &empty.Empty{})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(printStatus(status))
	return 0
}

func printStatus(status *proto.StatusResponse) string {
	printHeader := func(h *proto.Header) string {
		return formatKV([]string{
			fmt.Sprintf("Hash|%s", h.Hash),
			fmt.Sprintf("Number|%d", h.Number),
		})
	}
	full := []string{
		"Current Header",
		printHeader(status.CurrentHeader),
		"\nCurrent Block",
		printHeader(status.CurrentBlock),
	}
	return strings.Join(full, "\n")
}