package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

var from = time.Now().AddDate(-5, 0, 0)
var to = time.Now()

var syncCmd = cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Flags: []cli.Flag{
		&cli.TimestampFlag{
			Name:   "from",
			Layout: "2006-01-02",
		},
		&cli.TimestampFlag{
			Name:   "to",
			Layout: "2006-01-02",
		},
	},
	Usage: "TICKER",
	Action: func(c *cli.Context) error {
		if c.Timestamp("from") != nil {
			from = *c.Timestamp("from")
		}
		if c.Timestamp("to") != nil {
			to = *c.Timestamp("from")
		}
		return syncQuotes(c.Args().Slice(), from, to)
	},
}

var checkCmd = cli.Command{
	Name:    "check",
	Aliases: []string{"c"},
	Usage:   "check TICKER",
	Action: func(c *cli.Context) error {
		fmt.Println("completed task: ", c.Args().First())
		return nil
	},
}

var Commands = []*cli.Command{
	&checkCmd,
	&syncCmd,
}
