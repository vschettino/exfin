package main
import (
	"github.com/gin-gonic/gin"
	"github.com/vschettino/exfin/auth"
	r "github.com/vschettino/exfin/resources"
)
func router() (g *gin.Engine) {
	a := auth.JWTMiddleware()
	g = gin.Default()
	g.POST("/login", a.LoginHandler)
	authRequired := g.Group("/")
	authRequired.GET("/refresh_token", a.RefreshHandler)
	authRequired.Use(a.MiddlewareFunc())
	{
		authRequired.GET("/account", r.GetAccounts)
		authRequired.GET("/account/:id", r.GetAccount)
		authRequired.GET("/me", r.GetMyself)

	}

	return
}
