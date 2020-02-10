package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"marchito@gmail.com","password":"mdas"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}

	user, err := repository.UserLogin("marchito@gmail.com", "mdas")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid User Credentials", err.Message)
}
func TestLoginErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"marchito@gmail.com","password":"mdas"}`,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": "Interface error while trying to Login", "status": 500, "error": "internal_server_error"}`,
	})
	repository := usersRepository{}

	user, err := repository.UserLogin("marchito@gmail.com", "mdas")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Interface error while trying to Login", err.Message)
}

func TestLoginUserCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"marchito@gmail.com","password":"mdas"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid user credentials response", "status": 404, "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.UserLogin("marchito@gmail.com", "mdas")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Invalid user credentials response", err.Message)

}
func TestLoginInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"marchito@gmail.com","password":"mdas"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{ "id": "34",
    "first_name": "Marcio",
    "last_name": "DaSilva",
    "email": "marchito@gmail.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.UserLogin("marchito@gmail.com", "mdas")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error unmarshall user login response", err.Message)

}
func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"marchito@gmail.com","password":"mdas"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{ "id": 34,
"first_name": "Marcio",
"last_name": "DaSilva",
"email": "marchito@gmail.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.UserLogin("marchito@gmail.com", "mdas")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 34, user.Id)
	assert.EqualValues(t, "Marcio", user.FirstName)
	assert.EqualValues(t, "DaSilva", user.LastName)
	assert.EqualValues(t, "marchito@gmail.com", user.Email)
}
