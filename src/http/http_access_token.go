package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prosline/pl_oauth_api/src/domain/access_token"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (at *accessTokenHandler)GetById(c *gin.Context){
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := at.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
