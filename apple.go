package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	Apple struct {
	}
)

const appleKeysEndpoint = "https://appleid.apple.com/auth/keys"

func (s *Apple) auth(token string) (ud *UserDetails, err error) {
	return s.authWithCheckAUD(token, "")
}

func (s *Apple) authWithCheckAUD(token, aud string) (ud *UserDetails, err error) {
	type TokenInfo struct {
		Sub   string `json:"sub"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	t, err := new(jwt.Parser).ParseWithClaims(token, &TokenInfo{}, getTokenValidateFunc(appleKeysEndpoint))
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
			ID:    info.Sub,
			Email: info.Email,
		}
	}
	return
}
