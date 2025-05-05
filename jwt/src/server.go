package main

import (
	"jwt/routes"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	listenAddress string
	router        *gin.Engine
}

func NewApiServer(listenAddress string) *ApiServer {
	//TODO: Add logger to the Gin engine, so it uses JSON format also
	return &ApiServer{
		listenAddress: listenAddress,
		router:        gin.New(),
	}
}

func (apiServer *ApiServer) Start() error {
	apiServer.router.Use(gin.Recovery())
	apiServer.router.Use(gin.Logger())

	// Define routes
	routes.RegisterSignupRoute(apiServer.router, database, &logger)
	routes.RegisterLoginRoute(apiServer.router, database, &logger)
	routes.RegisterLogoutRoute(apiServer.router)
	routes.RegisterAPI1Route(apiServer.router, &logger)
	routes.RegisterAPI2Route(apiServer.router, database, &logger)

	// Log server start
	logger.Info().Msgf("Server started on address %s", apiServer.listenAddress)

	// Start the server
	return apiServer.router.Run(apiServer.listenAddress)
}
