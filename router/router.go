package router

import (
	"github.com/gin-gonic/gin"
	"gitlab/vote-crossbow/configs"
	"gitlab/vote-crossbow/features/registration_vote_crossbow_v1"
)

var repository = registration_vote_crossbow_v1.ProvideRepository(configs.ConnectDB())
var service = registration_vote_crossbow_v1.ProvideService(repository)
var controller = registration_vote_crossbow_v1.ProvideController(service)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		v1.GET("", controller.GetRegistrationVoteCrossbow)
		return r
	}

}
