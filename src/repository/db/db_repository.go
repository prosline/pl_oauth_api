package db

import (
	"github.com/gocql/gocql"
	"github.com/prosline/pl_logger/logger"
	"github.com/prosline/pl_oauth_api/src/clients/cassandra"
	"github.com/prosline/pl_oauth_api/src/domain/access_token"
	"github.com/prosline/pl_util/utils/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpiration  = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessToken) *rest_errors.RestErr
	UpdateExpiration(access_token.AccessToken) *rest_errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("Token not found")
		}
		return nil, rest_errors.NewInternalServerError("Error getting Id - Database Error",err)
	}
	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *rest_errors.RestErr {
	if er := cassandra.GetSession().Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires,
	).Exec(); er != nil {
		logger.Info(er.Error())
		rest_errors.NewInternalServerError("Error creating access token in Cassandra",er)
	}
	return nil
}
func (r *dbRepository) UpdateExpiration(token access_token.AccessToken) *rest_errors.RestErr {
	if er := cassandra.GetSession().Query(queryUpdateExpiration, token.Expires, token.AccessToken).Exec(); er != nil {
		rest_errors.NewInternalServerError("Error updating access token and expire information!",er)
	}
	return nil
}
