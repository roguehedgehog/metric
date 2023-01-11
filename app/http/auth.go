package http

import (
	"github.com/gin-gonic/gin"
	"github.com/roguehedgehog/metric/domain/user"
	"github.com/roguehedgehog/metric/infra"
	"github.com/rs/zerolog/log"
	"net/http"
)

type authReq struct {
	username string `form:"username" binding:"required"`
	password string `form:"password" binding:"required"`
}

func login(c *gin.Context) {
	svc := user.NewLoginSvc(infra.NewUserRepo(infra.PrimaryDb))
	var req authReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := svc.Login(req.username, req.password); err != nil {
		switch err.(type) {
		case user.NotAuthorisedError, user.NotFoundError:
			c.Writer.WriteHeader(http.StatusUnauthorized)
		default:
			log.Err(err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

}
