package router

import (
	"github.com/gin-gonic/gin"
	"gitlab/vote-crossbow/configs"
	"gitlab/vote-crossbow/features/registration_vote_crossbow_v1"
	"gitlab/vote-crossbow/features/user_vote_crossbow_v1"
)

var repository = registration_vote_crossbow_v1.ProvideRepository(configs.ConnectDB())
var service = registration_vote_crossbow_v1.ProvideService(repository)
var controller = registration_vote_crossbow_v1.ProvideController(service)

var repositoryUser = user_vote_crossbow_v1.ProvideRepository(configs.ConnectDB())
var serviceUser = user_vote_crossbow_v1.ProvideService(repositoryUser)
var controllerUser = user_vote_crossbow_v1.ProvideController(serviceUser)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		v1.GET("", controller.GetRegistrationVoteCrossbow)

		v1.POST("/user/registration", controllerUser.CreateVoteCrossbow)
		v1.POST("/user/login", controllerUser.LoginVoteCrossbow)
		return r
	}

}
