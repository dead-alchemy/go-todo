package main

import (
	"log"
	util "server/utils"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// set port.
	port := util.GoDotEnv("GO_PORT")

	//set up router.
	router := SetUp()

	// run router
	log.Fatal(router.Run(":" + port))
}

func SetUp() *gin.Engine {
	// init router
	router := gin.Default()

	// setup mode to run
	env := util.GoDotEnv("GO_ENV")
	if env != "production" && env != "test" {
		gin.SetMode(gin.DebugMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	router.Use(helmet.Default())

	// put all routes here.

	return router

}
