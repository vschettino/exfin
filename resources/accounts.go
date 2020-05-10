package resources

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/vschettino/exfin/db"
	m "github.com/vschettino/exfin/models"
	"net/http"
)

type CreateAccountRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=255"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

func (r CreateAccountRequest) ToAccount() m.Account {
	acc := m.Account{Name: r.Name, Email: r.Email}
	acc.SetHashPassword(r.Password)
	return acc
}

type UpdateAccountRequest struct {
	Name     string `json:"name" binding:"omitempty,min=3,max=255"`
	Email    string `json:"email" binding:"omitempty,email,max=255"`
	Password string `json:"password" binding:"omitempty,min=8,max=255"`
}

func (r UpdateAccountRequest) UpdateAccount(acc *m.Account) *m.Account {
	if r.Email != "" {
		acc.Email = r.Email
	}
	if r.Name != "" {
		acc.Name = r.Name
	}
	if r.Password != "" {
		acc.SetHashPassword(r.Password)
	}
	return acc
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
	var req FetchByIdUri
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

func CreateAccount(c *gin.Context) {
	var conn = db.Connection()
	var request CreateAccountRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	account := request.ToAccount()
	err := conn.Insert(&account)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "This email already exists"})
		return
	}
	c.JSON(http.StatusCreated, &account)
}

func UpdateAccount(c *gin.Context) {
	var conn = db.Connection()
	var request UpdateAccountRequest
	var uri FetchByIdUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var account = m.Account{Id: uri.Id}
	if err := conn.Select(&account); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found"})
		return
	}
	request.UpdateAccount(&account)
	if err := conn.Update(&account); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "This email already exists"})
		return
	}
	c.JSON(http.StatusOK, &account)
}
