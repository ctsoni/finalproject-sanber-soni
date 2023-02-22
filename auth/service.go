package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY = []byte("SeECre3t_K3y_T0k3n_f1n4L_PR0j3Ct")

// GenerateToken in function to generate JWT token with userID and isAdmin as payload
// and SECRET_KEY as its secret key
func GenerateToken(userID int, isAdmin bool) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	payload["isAdmin"] = isAdmin

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// ValidateToken in function to validate encodedToken JWT token
// and return validated token
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		encodedToken,
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(SECRET_KEY), nil
		})

	if err != nil {
		return token, err
	}

	return token, nil
}
