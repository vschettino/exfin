package resources_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	m "github.com/vschettino/exfin/models"
	"github.com/vschettino/exfin/tests"
	"testing"
)


func TestGetMyselfUnauthorized(t *testing.T) {
	w, _ := tests.MakeGET("/me", "")
	assert.Equal(t, 401, w.Code)
}

func TestGetMyselfSuccessful(t *testing.T) {
	tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakeGET("/me", tests.GenerateToken("admin@exfin.org", 1))
	var account m.Account
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &account))
	assert.Equal(t, uint(1), account.Id )
}

func TestAccountsUnauthorized(t *testing.T) {
	w, _ := tests.MakeGET("/accounts", "")
	assert.Equal(t, 401, w.Code)
}


func TestGetAccountsSuccessful(t *testing.T) {
	tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakeGET("/accounts", tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 200, w.Code)
}

func TestAccountUnauthorized(t *testing.T) {
	w, _ := tests.MakeGET("/accounts/1", "")
	assert.Equal(t, 401, w.Code)
}


func TestGetAccountDoesntExist(t *testing.T) {
	w, _ := tests.MakeGET("/accounts/666", tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 404, w.Code)
	w.Body.String()
}

func TestGetAccountBadParams(t *testing.T) {
	w, _ := tests.MakeGET("/accounts/stringid666", tests.GenerateToken("admin@exfin.org", 1))
	assert.Equal(t, 400, w.Code)
	w.Body.String()
}


func TestGetAccountSuccessful(t *testing.T) {
	w, _ := tests.MakeGET("/accounts/1", tests.GenerateToken("admin@exfin.org", 1))
	var account m.Account
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &account))
	assert.Equal(t, uint(1), account.Id )
}

