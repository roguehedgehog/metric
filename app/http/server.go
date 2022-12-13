package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.String(http.StatusOK, "I'm alive")
	})

	return r

}

func (s *Server) Start() {
	s.router.Run(fmt.Sprintf(":%d", s.port))
}
