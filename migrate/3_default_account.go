package main

import (
	"fmt"
	"github.com/go-pg/migrations/v7"

)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("make account email unique")
		_, err := db.Exec(`create unique index accounts_email_uindex on accounts (email);`)
		fmt.Println("add default admin user")
		_, err = db.Exec(`INSERT INTO "public"."accounts" ("id", "email", "password", "name") VALUES (DEFAULT, 'admin@exfin.org', '$2a$10$Yt5xbVhaU8D3KrsrJk6nGONoYz6x5gfDe9tmMrYyDTtvYZft1vA.a', 'Default Admin');`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping unique_index...")
		_, err := db.Exec(`drop index accounts_email_uindex; DELETE FROM "public"."accounts" WHERE "email" = 'admin@exfin.org'`)
		return err
	})
}
