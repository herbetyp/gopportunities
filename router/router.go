package router

import (
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	router := gin.Default()

	initRoutes(router)

	router.Run(":8080")
}
