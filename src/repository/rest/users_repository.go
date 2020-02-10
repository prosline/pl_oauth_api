package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/prosline/pl_oauth_api/src/domain/users"
	"github.com/prosline/pl_util/utils/rest_errors"
	"strings"
	"time"
)

var (
	userRequestClient = rest.RequestBuilder{
		Timeout: 100 * time.Millisecond,
		BaseURL: "http://localhost:8080",
	}
)

type RestUserRepository interface {
	UserLogin(string, string) (*users.User, *rest_errors.RestErr)
}
type usersRepository struct{}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (ur *usersRepository) UserLogin(login string, password string) (*users.User, *rest_errors.RestErr) {
	if (strings.TrimSpace(login) == "") || (strings.TrimSpace(password) == "") {
		err := rest_errors.NewBadRequestError("Invalid Login and Password")
		return nil, err
	}
	credentials := users.LoginRequest{
		Email:    login,
		Password: password,
	}
	resp := userRequestClient.Post("/users/login", credentials)

	if resp == nil || resp.Response == nil {
		return nil, rest_errors.NewInternalServerError("Invalid User Credentials",errors.New("Invalid http.Response"))
	}
	if resp.StatusCode > 299 {
		fmt.Println(resp.String())
		apiErr, err := rest_errors.NewRestErrorFromBytes(resp.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("Interface error while trying to Login", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(resp.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("Error unmarshall user login response",err)
	}
	return &user, nil
}
