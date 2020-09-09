package auth

import (
	"errors"
)

type (
	AuthType uint

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
	AuthTypeGoogle AuthType = iota + 1
	AuthTypeApple
	AuthTypeFacebook
	AuthTypeVK
)

func Auth(token string, authType AuthType) (userDetails *UserDetails, err error) {
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

	ud, err := authService.auth(token)
	if err != nil {
		return nil, err
	}
	return ud, nil
}
