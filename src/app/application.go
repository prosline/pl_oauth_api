package app

import (
	"github.com/gin-gonic/gin"
	"github.com/prosline/pl_logger/logger"
	"github.com/prosline/pl_oauth_api/src/domain/access_token"
	"github.com/prosline/pl_oauth_api/src/http"
	"github.com/prosline/pl_oauth_api/src/repository/db"
)
var (
	router = gin.Default()
)

func StartApplication() {
	srv := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(srv)
	router.GET("/oauth/access_token/:access_token_id",atHandler.GetById)
	logger.Info("Starting Application....")
	router.Run((":8000"))
}
