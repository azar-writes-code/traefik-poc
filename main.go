package main

import (
	"os"

	restcontrollers "github.com/azar-writes-code/traefik-poc/user-service/pkg/rest/server/controllers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	// rest server configuration
	router := gin.Default()

	userController, err := restcontrollers.NewUserController()
	if err != nil {
		log.Errorf("error occurred: %v", err)
		os.Exit(1)
	}

	v1 := router.Group("/v1")
	{

		v1.POST("/users", userController.CreateUser)

		v1.GET("/users/:id", userController.FetchUser)

		v1.PUT("/users/:id", userController.UpdateUser)

		v1.DELETE("/users/:id", userController.DeleteUser)

	}

	Port := ":9000"
	log.Println("Server started")
	if err = router.Run(Port); err != nil {
		log.Errorf("error occurred: %v", err)
		os.Exit(1)
	}

}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
