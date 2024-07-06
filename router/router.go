package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	initRoutes(router)

	router.Run(":8080")
}
