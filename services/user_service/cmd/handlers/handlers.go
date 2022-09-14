package handlers

import (
	"net/http"

	"user_service/module/user/usertransport/ginuser"

	goservice "github.com/phathdt/go-sdk"
	"github.com/phathdt/go-sdk/httpserver/middleware"

	"github.com/gin-gonic/gin"
)

func Router(sc goservice.ServiceContext) func(engine *gin.Engine) {
	return func(engine *gin.Engine) {
		engine.Use(middleware.Recover(sc))
		engine.Use(middleware.AllowCORS())

		engine.GET("/ping", ping)

		api := engine.Group("api")
		{
			users := api.Group("/users")
			{
				users.POST("/signup", ginuser.SignupUser(sc))
				users.POST("/login", ginuser.LoginUser(sc))
			}
		}
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
