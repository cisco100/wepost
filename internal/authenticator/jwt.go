package authenticator

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret   string
	audience string
	issue    string
	expiry   time.Duration
}

func NewJWTAuthenticator(secret, audience, issue string, expiry time.Duration) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret:   secret,
		audience: audience,
		issue:    issue,
		expiry:   expiry,
	}
}

func (ja *JWTAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ja.secret))
	if err != nil {
		return " ", err
	}
	return tokenString, nil
}

func (ja *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method used ,%s", t.Header["alg"])

		}
		return []byte(ja.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(ja.audience),
		jwt.WithIssuer(ja.issue),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
