package server

import (
	"github.com/Sudhanshu-NITR/notification-service/internal/routes"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	routes.Register(r)

	return r
}
