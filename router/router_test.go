package router_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vschettino/exfin/router"
	"testing"
)

func TestInitializeRouter(t *testing.T) {
	r := router.Router()
	assert.IsType(t, gin.Engine{}, *r )
}