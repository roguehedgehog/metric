package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	port   int
	router *gin.Engine
}

func NewServer(port int) Server {
	return Server{
		port:   port,
		router: NewRouter(),
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		app := HealthCheck()
		if app.isHealthy() {
			c.JSON(http.StatusOK, app)
			return
		}

		c.JSON(http.StatusServiceUnavailable, app)
	})
	r.POST("/users", newUser)

	return r

}

func (s *Server) Start() {
	s.router.Run(fmt.Sprintf(":%d", s.port))
}
