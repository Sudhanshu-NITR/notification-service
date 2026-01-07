package server

import (
	"github.com/Sudhanshu-NITR/notification-service/internal/config"
	"github.com/Sudhanshu-NITR/notification-service/internal/routes"
	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	routes.Register(r)

	return r
}
