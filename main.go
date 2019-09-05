package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zavier/front-api/api"
	"github.com/zavier/front-api/middleware"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// 允许跨越访问
	router.Use(middleware.Cors())

	api.ConfigHero(router)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal("server start error", err)
	}

}
