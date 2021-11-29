package auth

import "testing"

var (
	googleToken = "" // Add valid google token for test
	googleAudience = "816826508635-e93fbc53bj7bo8krv2gnoghbg7fi37d7.apps.googleusercontent.com" // Change token audience

	appleToken = "" // Add valid apple token for test
	appleAudience = "com.sleepeasy.sleep"  // Change token audience
)

func TestAuthWithAUD(t *testing.T) {
	type test struct {
        token string
        audience   string
        authType Type
		expectedValid bool
    }

	tests := []test{
		{
			token: googleToken,
			audience: googleAudience,
			authType: AuthTypeGoogle,
			expectedValid: true,
		},
		{
			token: googleToken,
			audience: "",
			authType: AuthTypeGoogle,
			expectedValid: true,
		},
		{
			token: googleToken,
			audience: googleAudience + "_",
			authType: AuthTypeGoogle,
			expectedValid: false,
		},
		{
			token: appleToken,
			audience: appleAudience,
			authType: AuthTypeApple,
			expectedValid: true,
		},
		{
			token: appleToken,
			audience: "",
			authType: AuthTypeApple,
			expectedValid: true,
		},
		{
			token: appleToken,
			audience: appleAudience + "_",
			authType: AuthTypeApple,
			expectedValid: false,
		},
	}

	for _, tc := range tests {
		ud, err := AuthWithCheckAUD(tc.token, tc.audience, tc.authType)
		if err != nil && tc.expectedValid {
			t.Fatalf("Unexpected invalid result (err)")
		}

		if ud == nil && tc.expectedValid {
			t.Fatalf("Unexpected invalid result (empty user data)")
		}

		if err == nil && !tc.expectedValid {
			t.Fatalf("Unexpected valid result (err is nil)")
		}
	}
}