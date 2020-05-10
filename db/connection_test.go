package db_test

import (
	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/db"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	conn := db.Connection()
	assert.Contains(t, conn.String(), "DB")
}

func TestQueryLogger(t *testing.T) {
	logger := db.DbLogger{}
	query := pg.QueryEvent{
		StartTime: time.Time{},
		DB:        db.Connection(),
	}
	c := context.Background()
	ctx, err := logger.BeforeQuery(c, &query)
	assert.NoError(t, err)
	assert.Equal(t, c, ctx)
	assert.NoError(t, logger.AfterQuery(c, &query))
}
