package db

import (
	"github.com/prosline/pl_oauth_api/src/domain/access_token"
	"github.com/prosline/pl_oauth_api/src/utils/errors"
)
func NewRepository() DbRepository {
	return &dbRepository{}
}
type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository)GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Database Connection not implemented")
}
