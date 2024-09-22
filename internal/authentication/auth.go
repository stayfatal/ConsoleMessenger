package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	secretKey  = []byte("super-secret-key")
	sighError  = errors.New("unknown sigh method")
	tokenError = errors.New("invalid token")
)

type Claims struct {
	Id int
	jwt.StandardClaims
}

func CreateToken(Id int) (string, error) {
	claims := &Claims{
		Id: Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secretKey)

	return t, errors.Wrap(err, "signing token")
}

func ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrap(sighError, "signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.Wrap(tokenError, "validating token")
	}

	return claims, nil
}
