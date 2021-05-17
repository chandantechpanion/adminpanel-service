package server

import (
	"fmt"

	"github.com/chandantechpanion/adminpanel-service/pkg/config"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/chandantechpanion/adminpanel-service/docs"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *dserver {
	return &dserver{
		router: e,
		cont:   c,
		logger: l,
	}
}

// Start start serving the application

func (ds *dserver) CorsConfig() {
	ds.router.Use(cors.Default())

}

func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}

	//url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	ds.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
