package models_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/models"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, _ := models.HashPassword("myneatpass")
	cost, _ := bcrypt.Cost([]byte(hash))
	assert.Equal(t, bcrypt.DefaultCost, cost)
}

func TestVerifyPassword(t *testing.T) {
	password, err := models.HashPassword("myneatpass")
	var account = models.Account{
		Password: password,
	}
	assert.NoError(t, err)
	assert.True(t, account.VerifyPassword("myneatpass"))
}

func TestAccountToString(t *testing.T) {
	acc := models.Account{
		Id: 6644,
	}
	assert.Contains(t, acc.String(), "6644")
}
