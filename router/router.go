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
		authRequired.GET("/accounts", r.GetAccounts)
		authRequired.POST("/accounts", r.CreateAccount)
		authRequired.GET("/accounts/:id", r.GetAccount)
		authRequired.PATCH("/accounts/:id", r.UpdateAccount)
		authRequired.GET("/me", r.GetMyself)

	}

	return
}
