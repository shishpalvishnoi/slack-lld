package http_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"slack-application/log"
	"sync"
	"syscall"
	"time"
)

var logger = log.GetNewLogger()

func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()

	router := gin.New()
	router.Use(gin.Recovery())

	// initialize routes
	initializeRoutes(router)

	httpServer := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// start server in separate goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error("Error starting server")
		}
	}()

	// make channel quit for getting signal for shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")
}

func initializeRoutes(router *gin.Engine) {
	router.GET("/health", healthCheck)

	slackRoutes := router.Group("/slack")
	{
		slackRoutes.POST("/registerUser", registerUser)
		slackRoutes.POST("/createWorkspace", createWorkspace)
		slackRoutes.POST("/createChannelInWorkspace", createChannelInWorkspace)
		slackRoutes.POST("/addUserToWorkspace", addUserToWorkspace)
		slackRoutes.POST("/addUserToChannel", addUserToChannel)
		slackRoutes.POST("/sendMessageToChannel", sendMessageToChannel)
		slackRoutes.POST("/sendMessageToUser", sendMessageToUser)
	}

}
