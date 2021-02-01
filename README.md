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
import "github.com/rosberry/auth"

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

### Important
- VK don't sent email in user information
- Apple don't include username (firstname and lastname) in JWT token

## About

<img src="https://github.com/rosberry/Foundation/blob/master/Assets/full_logo.png?raw=true" height="100" />

This project is owned and maintained by [Rosberry](http://rosberry.com). We build mobile apps for users worldwide 🌏.

Check out our [open source projects](https://github.com/rosberry), read [our blog](https://medium.com/@Rosberry) or give us a high-five on 🐦 [@rosberryapps](http://twitter.com/RosberryApps).

## License

This project is available under the MIT license. See the LICENSE file for more info.
