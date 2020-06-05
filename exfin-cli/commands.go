package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/markcheno/go-quote"
	"github.com/urfave/cli/v2"
	"github.com/vschettino/exfin/db"
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
		api := db.InfluxWriteApi()
		for lineNumber := range spy.Close {
			p := influxdb2.NewPoint("stocks", map[string]string{
				"ticker": spy.Symbol,
			}, map[string]interface{}{
				"close":  spy.Close[lineNumber],
				"higher": spy.High[lineNumber],
			}, spy.Date[lineNumber])
			api.WritePoint(p)
		}
		api.Flush()
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
