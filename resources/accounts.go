package resources

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/vschettino/exfin/db"
	m "github.com/vschettino/exfin/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AccountRequest struct {
	Id uint `uri:"id" binding:"required"`
}

func GetAccounts(c *gin.Context) {
	var accounts []m.Account
	var conn = db.Connection()
	err := conn.Model(&accounts).Select()
	if err != nil {
		panic(err)
	}
	c.JSON(200, accounts)
}
func GetAccount(c *gin.Context) {
	var req AccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
		return
	}
	var conn = db.Connection()
	var account = m.Account{Id: req.Id}
	if err := conn.Select(&account); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found"})
		return
	}
	c.JSON(200, account)
}

func GetMyself(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, claims)
}

func GeneratePass(c *gin.Context) {
	hash, err := bcrypt.GenerateFromPassword([]byte("adminpass"), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(hash))
	c.JSON(200, "")
}
