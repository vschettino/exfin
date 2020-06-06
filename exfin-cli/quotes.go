package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/markcheno/go-quote"
	"github.com/vschettino/exfin/db"
	"sync"
	"time"
)

func syncQuotes(tickers []string, from time.Time, to time.Time) error {
	var group sync.WaitGroup
	group.Add(len(tickers))
	for _, ticker := range tickers {
		go writeQuoteIntoInflux(ticker, from, to, &group)
	}
	group.Wait()
	return nil
}

func writeQuoteIntoInflux(ticker string, from time.Time, to time.Time, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("[%s] Start Fetch from  %s to %s\n", ticker, from, to)
	api := db.InfluxWriteApi()
	quotes, err := quote.NewQuoteFromYahoo(
		ticker,
		from.Format("2006-01-02"),
		to.Format("2006-01-02"),
		quote.Daily, true)
	if err != nil {
		fmt.Printf("[%s] Error fetching: %s\n", ticker, err)
		return
	}
	totalDataPoints := len(quotes.Close)
	fmt.Printf("[%s] Fetched %d data points\n", ticker, totalDataPoints)
	for lineNumber := range quotes.Close {
		p := influxdb2.NewPoint("stocks", map[string]string{
			"ticker": quotes.Symbol,
		}, map[string]interface{}{
			"close":  quotes.Close[lineNumber],
			"higher": quotes.High[lineNumber],
			"volume": quotes.Volume[lineNumber],
		}, quotes.Date[lineNumber])
		api.WritePoint(p)
	}
	api.Flush()
	fmt.Printf("[%s] Done!\n", ticker)
	return
}
