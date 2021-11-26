package auth

import "testing"

// Add tokens for test
var googleToken = "" 
var appleToken = ""

func TestAuthWithBadTime(t *testing.T) {
	// Leeway = 60 * 60 * 4 // You can change leeway
	if googleToken != "" {
		ud, err := Auth(googleToken, AuthTypeGoogle)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		if ud == nil {
			t.Fatalf("user data is nil")
		}
	}

	if appleToken != "" {
		ud, err := Auth(appleToken, AuthTypeApple)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		if ud == nil {
			t.Fatalf("user data is nil")
		}
	}
}