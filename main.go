package main

import (
	"github.com/chandantechpanion/adminpanel-service/cmd/server"
	"github.com/chandantechpanion/adminpanel-service/pkg/di"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// @title Techpanion Admin Portal API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.techpanion.com/support
// @contact.email chandresh@techpanion.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl /access_token/

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl /access_token/

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /access_token/

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl /access_token/
// @authorizationUrl /access_token/
// @scope.admin Grants read and write access to administrative information

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	g := gin.Default()
	d := di.BuildContainer()

	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})

	svr := server.NewServer(g, d, l)
	svr.CorsConfig()
	svr.MapRoutes()

	return svr.Start()
}
