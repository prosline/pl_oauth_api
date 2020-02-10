package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
func TestExpirationTime (t *testing.T){
	assert.EqualValues(t, 24, expirationTime, "Time should be defined based on 24 hours!")
}

func TestGetNewAccessToken(t *testing.T) {
	var userid int64
	userid = 34

	at := GetNewAccessToken(userid)
	assert.False(t, at.IsExpired(), "New access token should not be expired")
	assert.EqualValues(t,"", at.AccessToken, "New access token should not have defined access token id")
	assert.EqualValues(t,0, at.UserId, "New access token must not have User Id")
	assert.EqualValues(t,0, at.ClientId, "New access token must not have associated Client Id!")
	assert.EqualValues(t,userid, at.UserId, "New access token must have associated User Id!")
}
func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t,at.IsExpired(),"Empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(),"Access token setup for 3 hours from now should not be expired!")
}
