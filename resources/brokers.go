package resources

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/vschettino/exfin/db"
	m "github.com/vschettino/exfin/models"
	"net/http"
)

type CreateBrokerRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=255"`
	Institution string `json:"institution" binding:"required,min=3,max=255"`
}

func (r CreateBrokerRequest) ToBroker() m.Broker {
	bk := m.Broker{Name: r.Name, Institution: r.Institution}
	return bk
}

type UpdateBrokerRequest struct {
	Name        string `json:"name" binding:"omitempty,min=3,max=255"`
	Institution string `json:"institution" binding:"omitempty,min=3,max=255"`
}

func (r UpdateBrokerRequest) UpdateBroker(bk *m.Broker) *m.Broker {
	if r.Name != "" {
		bk.Name = r.Name
	}
	if r.Institution != "" {
		bk.Institution = r.Institution
	}
	return bk
}

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
	var req FetchByIdUri
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

func CreateBroker(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	var conn = db.Connection()
	var request CreateBrokerRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	broker := request.ToBroker()
	broker.AccountId = uint(claims["Id"].(float64))
	err := conn.Insert(&broker)
	_ = conn.Model(&broker).Relation("Account").Select()
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusCreated, &broker)
}

func UpdateBroker(c *gin.Context) {
	var conn = db.Connection()
	var request UpdateBrokerRequest
	var uri FetchByIdUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var broker = m.Broker{Id: uri.Id}
	if err := conn.Select(&broker); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found"})
		return
	}
	request.UpdateBroker(&broker)
	if err := conn.Update(&broker); err != nil {
		panic(err)
		return
	}
	_ = conn.Model(&broker).Relation("Account").Select()
	c.JSON(http.StatusOK, &broker)
}
