package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	Google struct {
	}
)

const googleKeysEndpoint = "https://www.googleapis.com/oauth2/v2/certs"

func (s *Google) auth(token string) (ud *UserDetails, err error) {
	return s.authWithCheckAUD(token, "")
}

func (s *Google) authWithCheckAUD(token, aud string) (ud *UserDetails, err error) {
	type TokenInfo struct {
		Sub       string `json:"sub"`
		Name      string `json:"name"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Email     string `json:"email"`
		Picture   string `json:"picture"`
		jwt.StandardClaims
	}

	t, err := new(jwt.Parser).ParseWithClaims(token, &TokenInfo{}, getTokenValidateFunc(googleKeysEndpoint))
	if err != nil {
		return nil, err
	}
	info, _ := t.Claims.(*TokenInfo)

	if aud != "" {
		if ok := info.VerifyAudience(aud, true); !ok {
			return nil, ErrNotValidAudience
		}
	}

	if info != nil {
		ud = &UserDetails{
			ID:        info.Sub,
			UserName:  info.Name,
			FirstName: info.FirstName,
			LastName:  info.LastName,
			Email:     info.Email,
			Picture:   info.Picture,
		}
	}
	return
}