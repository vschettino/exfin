package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/tests"
	"log"
	"testing"
)

func TestLoginSuccessful(t *testing.T) {
	w, _ := tests.MakePOST("/login", tests.RequestLogin())
	log.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
