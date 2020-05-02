package auth_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/tests"
	"testing"
)

func LoginRequest() map[string]string{
	return map[string]string{
		"username": "admin@exfin.org",
		"password": "adminpass",
	}
}


func TestLoginSuccessful(t *testing.T) {
	w, _ := tests.MakePOST("/login", LoginRequest())
	assert.Equal(t, 200, w.Code)
}

func TestLoginBadCredentials(t *testing.T) {
	login := LoginRequest()
	login["password"] = "thisbadauth"
	w, _ := tests.MakePOST("/login", login)
	assert.Equal(t, 401, w.Code)
}
