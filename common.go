package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

type (
	Claim struct {
		jwt.StandardClaims
	}

	AuthService interface {
		auth(token string) (ud *UserDetails, err error)
		authWithCheckAUD(token, aud string) (ud *UserDetails, err error)
	}
)

var ErrNotValidAudience = errors.New("not valid audience")

func getTokenValidateFunc(endpoint string) func(t *jwt.Token) (interface{}, error) {
	return func(t *jwt.Token) (interface{}, error) {
		// get the public key for verifying the identity token signature
		kid := t.Header["kid"].(string)
		set, err := jwk.FetchHTTP(endpoint, jwk.WithHTTPClient(&http.Client{}))
		if err != nil {
			return nil, err
		}
		selectedKey := set.Keys[0]
		for _, key := range set.Keys {
			if key.KeyID() == kid {
				selectedKey = key
				break
			}
		}
		var pubKeyIface interface{}
		selectedKey.Raw(&pubKeyIface)
		pubKey, ok := pubKeyIface.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf(`expected RSA public key from %s`, endpoint)
		}
		return pubKey, nil
	}
}

func jwtValidate(token string, out *Claim, endpoint string) (*jwt.Token, error) {
	t, err := new(jwt.Parser).ParseWithClaims(token, out, getTokenValidateFunc(endpoint))
	if err != nil {
		return nil, err
	}
	return t, nil
}
