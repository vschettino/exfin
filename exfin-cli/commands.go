package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/urfave/cli/v2"
	"time"
)

var from = time.Now().AddDate(-12, 0, 0).Format("2006-01-02")
var to = time.Now().Format("2006-01-02")

var sync = cli.Command{
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
		ticker := c.Args().First()
		if c.Timestamp("from") != nil {
			from = c.Timestamp("from").Format("2006-01-02")
		}
		if c.Timestamp("to") != nil {
			to = c.Timestamp("from").Format("2006-01-02")
		}
		spy, _ := quote.NewQuoteFromYahoo(ticker, from, to, quote.Daily, true)
		fmt.Print(spy.Close)
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
