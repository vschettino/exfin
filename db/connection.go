package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
	"os"
)

func Connection() *pg.DB {
	conn := pg.Connect(config())
	if os.Getenv("ENVIRONMENT") == "DEV" {
		conn.AddQueryHook(DbLogger{})
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

type DbLogger struct{}

func (d DbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	fmt.Println(q.Err)
	return nil
}
