package db_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/db"
	"testing"
)

func TestConnection(t *testing.T) {
	conn := db.Connection()
	assert.Contains(t, conn.String(), "DB")
}