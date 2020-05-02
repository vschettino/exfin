package auth_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/tests"
	"log"
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
	log.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
