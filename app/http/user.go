package http

import (
	"github.com/gin-gonic/gin"
	"github.com/roguehedgehog/metric/domain/user"
	"github.com/roguehedgehog/metric/infra"
	"github.com/rs/zerolog/log"
	"net/http"
)

type newUserReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func newUser(c *gin.Context) {
	s := user.NewCreateSvc(infra.NewUserRepo(infra.PrimaryDb))
	var req newUserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err := s.Create(req.Username, req.Password)
	if err == nil {
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	log.Err(err)
	switch t := err.(type) {
	case user.ExistsError:
		c.JSON(http.StatusBadRequest, gin.H{"errors": t.Error()})
	default:
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}
