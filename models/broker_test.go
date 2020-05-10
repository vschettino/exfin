package models_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/models"
	"testing"
)

func TestBrokerToString(t *testing.T) {
	acc := models.Broker{
		Id: 6644,
	}
	assert.Contains(t, acc.String(), "6644")
}
