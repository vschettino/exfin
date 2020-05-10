package resources

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/vschettino/exfin/db"
	m "github.com/vschettino/exfin/models"
	"net/http"
)

func GetBrokers(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var brokers []m.Broker
	var conn = db.Connection()
	err := conn.Model(&brokers).Relation("Account").Where("account_id = ?", claims["Id"]).Select()
	if err != nil {
		panic(err)
	}
	c.JSON(200, brokers)
}
func GetBroker(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var req FetchAccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
		return
	}
	var conn = db.Connection()
	var broker = m.Broker{}
	err := conn.Model(&broker).
		Relation("Account").
		Where("account_id = ?", claims["Id"]).
		Where("broker.id = ?", req.Id).
		Select()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found"})
		return
	}
	c.JSON(200, broker)
}
