package access_token

import (
	"github.com/prosline/pl_oauth_api/src/utils/errors"
	"strings"
)
type Repository interface{
	GetById(string) (*AccessToken, *errors.RestErr)
}
type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{ repository: repo,
	}
}

func (s *service)GetById(tokenId string) (*AccessToken, *errors.RestErr){
	if len(strings.TrimSpace(tokenId)) == 0 {
		return nil, errors.NewInternalServerError("Invalid access token!")
	}
	accessToken, err := s.repository.GetById(tokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
