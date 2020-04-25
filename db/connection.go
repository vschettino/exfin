package db

import (
	"github.com/go-pg/pg/v9"
	"os"
)

func Connection() *pg.DB{
	return pg.Connect(config())
}

func config() *pg.Options{
	return &pg.Options{
		Addr:     os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}