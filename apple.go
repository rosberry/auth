package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	Apple struct {
	}
)

type AppleTokenInfo struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (c *AppleTokenInfo) Valid() error {
	// leeway in seconds
	expiresLeeway := Leeway
	issuedLeeway := Leeway

	c.StandardClaims.ExpiresAt += expiresLeeway
	c.StandardClaims.IssuedAt -= issuedLeeway
	
	err := c.StandardClaims.Valid()
	
	c.StandardClaims.ExpiresAt -= expiresLeeway
	c.StandardClaims.IssuedAt += issuedLeeway
	
	return err
 }


const appleKeysEndpoint = "https://appleid.apple.com/auth/keys"

func (s *Apple) auth(token string) (ud *UserDetails, err error) {
	return s.authWithCheckAUD(token, "")
}

func (s *Apple) authWithCheckAUD(token, aud string) (ud *UserDetails, err error) {
	t, err := new(jwt.Parser).ParseWithClaims(token, &AppleTokenInfo{}, getTokenValidateFunc(appleKeysEndpoint))
	if err != nil {
		return nil, err
	}
	info, _ := t.Claims.(*AppleTokenInfo)

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
