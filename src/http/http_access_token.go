package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prosline/pl_oauth_api/src/domain/access_token"
	access_token2 "github.com/prosline/pl_oauth_api/src/services/access_token"
	"github.com/prosline/pl_util/utils/rest_errors"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpiration(*gin.Context)
}

type accessTokenHandler struct {
	service access_token2.Service
}

func NewHandler(service access_token2.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var atr access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		br := rest_errors.BadRequestError("Invalid JSON body")
		c.JSON(http.StatusBadRequest, br)
		return
	}
//	if err := atr.Validate(); err != nil {
//		er := errors.BadRequestError(err.Message)
//		c.JSON(er.Status, er.Message)
//		return
//	}
	accessToken, er := handler.service.Create(atr)
	if er != nil {
		c.JSON(er.Status(), er.Message)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
func (handler *accessTokenHandler) UpdateExpiration(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		br := rest_errors.BadRequestError("Invalid JSON body")
		c.JSON(http.StatusBadRequest, br)
		return
	}
	if err := at.Validate(); err != nil {
		er := rest_errors.BadRequestError(err.Message())
		c.JSON(http.StatusBadRequest, er.Message)
		return
	}
	if er := handler.service.UpdateExpiration(at); er != nil {
		c.JSON(http.StatusBadRequest, er.Message)
		return
	}
	c.JSON(http.StatusBadRequest, at)
}
