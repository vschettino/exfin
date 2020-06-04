package main

import (
	"flag"
	"fmt"
	"github.com/go-pg/migrations/v7"
	"github.com/vschettino/exfin/db"
	"os"
)

const usageText = `This program runs command on the db. Supported exfin-cli are:
  - init - creates version info table in the database
  - up - runs all available migrate.
  - up [target] - runs available migrate up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrate.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrate.
Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	conn := db.Connection()

	oldVersion, newVersion, err := migrations.Run(conn, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
