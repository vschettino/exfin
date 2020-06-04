package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

var today = cli.NewTimestamp(time.Now().AddDate(12, 0, 0))
var lastYear = cli.NewTimestamp(time.Now().AddDate(12, 0, 0))

var sync = cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Flags: []cli.Flag{
		&cli.TimestampFlag{
			Name:        "from",
			Usage:       "Fetch data from",
			Value:       lastYear,
			DefaultText: "One Year Ago",
			Layout:      "YYYY-MM-DD",
		},
		&cli.TimestampFlag{
			Name:        "to",
			Usage:       "Fetch data until",
			Value:       today,
			DefaultText: "Today",
			Layout:      "YYYY-MM-DD",
		},
	},
	Usage: "sync TICKER",
	Action: func(c *cli.Context) error {
		ticker := c.Args().First()
		fmt.Println("added task: ", c.Args().First())
		return nil
	},
}

var check = cli.Command{
	Name:    "check",
	Aliases: []string{"c"},
	Usage:   "check TICKER",
	Action: func(c *cli.Context) error {
		fmt.Println("completed task: ", c.Args().First())
		return nil
	},
}

var Commands = []*cli.Command{
	&check,
	&sync,
}
