package access_token

import (
	"errors"
	"github.com/prosline/pl_oauth_api/src/domain/access_token"
	"github.com/prosline/pl_oauth_api/src/repository/db"
	"github.com/prosline/pl_oauth_api/src/repository/rest"
	"github.com/prosline/pl_util/utils/rest_errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr)
	UpdateExpiration(access_token.AccessToken) *rest_errors.RestErr
}

type service struct {
	restUserRepo rest.RestUserRepository
	dbRepo db.DbRepository
}

func NewService(userRepo rest.RestUserRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUserRepo: userRepo,
		dbRepo: dbRepo,
	}
}

func (s *service)GetById(tokenId string) (*access_token.AccessToken, *rest_errors.RestErr){
	if len(strings.TrimSpace(tokenId)) == 0 {
		return nil, rest_errors.NewInternalServerError("Invalid access token!",errors.New("Invalid Token"))
	}
	accessToken, err := s.dbRepo.GetById(tokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create (tokenRequest access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr) {
//	if er := tokenRequest.Validate(); er != nil {
//		rest_errors.NewInternalServerError("Token id create errors! It can not be Zero")
//	}
	user, err := s.restUserRepo.UserLogin(tokenRequest.Username, tokenRequest.Password)
	if err != nil {
		return nil, err
	}
	// Generate brand new access token
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpiration (token access_token.AccessToken) *rest_errors.RestErr {
	if er := token.Validate(); er != nil {
		return rest_errors.NewInternalServerError("Token id update error! Token Id my be greather than ZERO",errors.New("Invalid Token"))
	}
	return s.dbRepo.UpdateExpiration(token)
}
