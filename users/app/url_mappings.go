package app

import (
	"github.com/Yapcheekian/microservices-go/users/controllers/ping"
	"github.com/Yapcheekian/microservices-go/users/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
