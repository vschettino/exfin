package main

import (
	"fmt"
	"github.com/go-pg/migrations/v7"

)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table account...")
		_, err := db.Exec(`CREATE TABLE accounts(
								id serial PRIMARY KEY,
							    email VARCHAR NOT NULL,
						        password VARCHAR,
								name VARCHAR NOT NULL )`,
		)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table account...")
		_, err := db.Exec(`DROP TABLE account`)
		return err
	})
}
