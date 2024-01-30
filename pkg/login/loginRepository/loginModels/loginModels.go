package loginModels

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	PrivateKey []byte
	PublicKey  []byte
}

var JwtToken JWT

// User represents a generic user
type UserRoleType string

const (
	Principal UserRoleType = "principal"
	Teacher   UserRoleType = "teacher"
	Student   UserRoleType = "student"
)

type User struct {
	ID       string       `json:"id" pg:",pk"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Role     UserRoleType `json:"role"` // "principal", "teacher", or "student"
}

func (j JWT) Create(content interface{}) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()
	expirationTime := time.Now().Add(time.Minute * 10)
	claims := make(jwt.MapClaims)
	claims["dat"] = content        // Our custom data.
	claims["exp"] = expirationTime // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()     // The time at which the token was issued.
	claims["nbf"] = now.Unix()     // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}
func (j JWT) Validate(token string) (interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.PublicKey)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["dat"], nil
}
