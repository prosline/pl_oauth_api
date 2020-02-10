package access_token

import (
	"errors"
	"fmt"
	"github.com/prosline/pl_util/utils/crypto"
	"github.com/prosline/pl_util/utils/rest_errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atr *AccessTokenRequest) Validate() *rest_errors.RestErr {
	switch atr.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return rest_errors.NewInternalServerError("Bad access token request",errors.New("Invalid Access Token"))
	}
	return nil
}

func (token AccessToken) Validate() *rest_errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if len(token.AccessToken) == 0 {
		return rest_errors.NewInternalServerError("Token id errors! It can not be Zero", errors.New("Access Token can not be Zero"))
	}
	if token.Expires <= 0 {
		return rest_errors.NewInternalServerError("Token expiration errors! It can not be Zero or negative number",errors.New("Expire can not be less than Zero"))
	}
	if token.ClientId <= 0 {
		return rest_errors.NewInternalServerError("Token client Id error! It can not be Zero or negative number", errors.New("Client Id can not be Zero"))
	}
	if token.UserId <= 0 {
		return rest_errors.NewInternalServerError("Token User Id error! It can not be Zero or negative number", errors.New("User Id can not be Zero"))
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token *AccessToken) IsExpired() bool {
	return time.Unix(token.Expires, 0).Before(time.Now().UTC())
}

func (token *AccessToken) Generate() {
	token.AccessToken = crypto.GetMd5(fmt.Sprintf("at-%d-%d-ran", token.UserId, token.Expires))
}
