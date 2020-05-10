package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vschettino/exfin/auth"
	r "github.com/vschettino/exfin/resources"
)

func Router() (g *gin.Engine) {
	a := auth.JWTMiddleware()
	g = gin.Default()
	g.POST("/login", a.LoginHandler)

	authRequired := g.Group("/")
	authRequired.GET("/refresh_token", a.RefreshHandler)
	authRequired.Use(a.MiddlewareFunc())
	{
		{ // Accounts
			authRequired.GET("/accounts", r.GetAccounts)
			authRequired.GET("/accounts/:id", r.GetAccount)
			authRequired.GET("/me", r.GetMyself)
			authRequired.POST("/accounts", r.CreateAccount)
			authRequired.PATCH("/accounts/:id", r.UpdateAccount)
		}

		{ // Brokers
			authRequired.GET("/brokers", r.GetBrokers)
			authRequired.GET("/brokers/:id", r.GetBroker)
			authRequired.POST("/brokers", r.CreateBroker)
			authRequired.PATCH("/brokers/:id", r.UpdateBroker)
		}
	}

	return
}
