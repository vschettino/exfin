package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"os"
)

func Connection() *pg.DB {
	conn := pg.Connect(config())
	if os.Getenv("ENVIRONMENT") == "DEV" {
		conn.AddQueryHook(Logger{})
	}
	return conn
}

func config() *pg.Options {
	return &pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

type Logger struct{}

func (d Logger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d Logger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	fmt.Println(q.Err)
	return nil
}

func InfluxClient() influxdb2.Client {
	return influxdb2.NewClientWithOptions("http://influx:8086", "",
		influxdb2.DefaultOptions().SetBatchSize(500))
}

func InfluxWriteApi() api.WriteApi {
	client := InfluxClient()
	return client.WriteApi("exfin", "exfin")
}
