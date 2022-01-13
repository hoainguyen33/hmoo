package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

var (
	jwtWraperSign = &JwtWrapper{
		SecretKey:       os.Getenv("SECRET_KEY"),
		Issuer:          os.Getenv("AUTH_SERVICE"),
		ExpirationHours: 24,
	}

	jwtWraperVerify = &JwtWrapper{
		SecretKey: os.Getenv("SECRET_KEY"),
		Issuer:    os.Getenv("AUTH_SERVICE"),
	}
)

func Sign(email string) (string, error) {
	token, err := jwtWraperSign.GenerateToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func Verify(token string) (*JwtClaim, error) {
	claims, err := jwtWraperVerify.ValidateToken(token)
	if err != nil {
		return &JwtClaim{}, err
	}
	return claims, nil
}

// GenerateToken generates a jwt token
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

//ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return nil, err
	}

	return claims, nil

}
