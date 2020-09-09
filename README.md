# Auth

Authentication with social networks, information about the user

## Supported:
- Apple
- Google
- Facebook
- VK


## Usage
### Input: token and auth type
```golang
ud, err := auth.Auth(googleToken, auth.AuthTypeGoogle)
```
Apple, Google: ```JWT token (id_token)```
Facebook, VK: ```access_token```

### Result: user details
```golang
UserDetails struct {
		ID        string
		FirstName string
		LastName  string
		UserName  string
		Email     string
		Picture   string
    }
```

## Important
- VK don't sent email in user information
- Apple don't include username (firstname and lastname) in JWT token