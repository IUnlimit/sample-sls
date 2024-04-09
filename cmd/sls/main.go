package sls

import (
	"fmt"
	global "github.com/IUnlimit/sample-sls/internal"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ServeRESTFul() {
	config := global.Config.API
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}

	err := r.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Errorf("API server exit with error: %v", err)
	}
}
