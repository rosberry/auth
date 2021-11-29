package auth

import (
	"errors"
)

type (
	//Type - social network
	Type uint

	//UserDetails - information about user from social network
	UserDetails struct {
		ID        string
		FirstName string
		LastName  string
		UserName  string
		Email     string
		Picture   string
	}
)

const (
	//AuthTypeGoogle - auth with Google
	AuthTypeGoogle Type = iota + 1
	//AuthTypeApple - auth with Apple
	AuthTypeApple
	//AuthTypeFacebook - auth with Facebook
	AuthTypeFacebook
	//AuthTypeVK - auth with VK
	AuthTypeVK
)

// Leeway in seconds (for issuedAt and ExpiresAt)
var Leeway int64 = 10

//Auth returning user details by token and auth type
func Auth(token string, authType Type) (userDetails *UserDetails, err error) {
	return auth(token, "", authType)
}

func AuthWithCheckAUD(token string, aud string, authType Type) (userDetails *UserDetails, err error) {
	return auth(token, aud, authType)
}

func auth(token string, aud string, authType Type) (userDetails *UserDetails, err error) {
	var authService AuthService
	switch authType {
	case AuthTypeGoogle:
		authService = &Google{}
	case AuthTypeApple:
		authService = &Apple{}
	case AuthTypeFacebook:
		authService = &Facebook{}
	case AuthTypeVK:
		authService = &VK{}
	default:
		authService = nil
	}

	if authService == nil {
		return nil, errors.New("Auth service is nil")
	}

	switch aud {
	case "": 
		userDetails, err = authService.auth(token)
	default:
		userDetails, err = authService.authWithCheckAUD(token, aud)
	}
	
	if err != nil {
		return nil, err
	}

	return userDetails, nil
}