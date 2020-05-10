package resources_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/db"
	m "github.com/vschettino/exfin/models"
	"github.com/vschettino/exfin/tests"
	"testing"
)

func CreateAccountRequest() map[string]string {
	return map[string]string{
		"email":    "newuser@exfin.org",
		"password": "newpass123",
		"name":     "New user",
	}
}

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
	assert.Equal(t, uint(1), account.Id)
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
	assert.Equal(t, uint(1), account.Id)
}

func TestCreateAccountMissingRequiredParameters(t *testing.T) {
	req := CreateAccountRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	delete(req, "name")
	w, _ := tests.MakePOST("/accounts", req, token)
	assert.Equal(t, 400, w.Code)
}

func TestCreateAccountWeakPassword(t *testing.T) {
	req := CreateAccountRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	req["password"] = "dumb"
	w, _ := tests.MakePOST("/accounts", req, token)
	assert.Equal(t, 400, w.Code)
}

func TestCreateAccountEmailAlreadyExists(t *testing.T) {
	req := CreateAccountRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	req["email"] = "admin@exfin.org"
	w, _ := tests.MakePOST("/accounts", req, token)
	assert.Equal(t, 422, w.Code)
}

func TestCreateAccountSuccessful(t *testing.T) {
	conn := db.Connection()
	var account m.Account
	req := CreateAccountRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePOST("/accounts", req, token)
	assert.Equal(t, 201, w.Code)
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &account))
	assert.Equal(t, account.Name, req["name"])
	t.Cleanup(func() {
		_ = conn.Delete(&account)
	})
}

func TestUpdateUnknownAccount(t *testing.T) {
	req := CreateAccountRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePATCH("/accounts/987987", req, token)
	assert.Equal(t, 404, w.Code)
}

func TestUpdateInvalidId(t *testing.T) {
	req := CreateAccountRequest()
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePATCH("/accounts/fid", req, token)
	assert.Equal(t, 400, w.Code)
}

func TestUpdateAccountEmailAlreadyExists(t *testing.T) {
	conn := db.Connection()
	req := CreateAccountRequest()
	account := m.Account{
		Name:  "My Name",
		Email: "emailtopatch@exfing.org",
	}
	_ = conn.Insert(&account)
	token := tests.GenerateToken("admin@exfin.org", 1)
	req["email"] = "admin@exfin.org"
	w, _ := tests.MakePATCH(fmt.Sprintf("/accounts/%d", account.Id), req, token)
	assert.Equal(t, 422, w.Code)
	t.Cleanup(func() {
		_ = conn.Delete(&account)
	})
}

func TestUpdateAccountSuccessful(t *testing.T) {
	conn := db.Connection()
	account := m.Account{
		Name:  "My Name",
		Email: "emailtopatch@exfing.org",
	}
	_ = conn.Insert(&account)
	req := CreateAccountRequest()
	delete(req, "password")
	token := tests.GenerateToken("admin@exfin.org", 1)
	w, _ := tests.MakePATCH(fmt.Sprintf("/accounts/%d", account.Id), req, token)
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal([]byte(w.Body.String()), &account))
	assert.Equal(t, account.Name, req["name"])
	assert.Equal(t, account.Name, req["name"])
	t.Cleanup(func() {
		_ = conn.Delete(&account)
	})
}
