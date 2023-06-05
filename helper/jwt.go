package helper

import (
	"crypto"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
}

type JWTCustomClaims struct {
	jwt.StandardClaims
}

func NewJWT(kmsClient *KMSClient) *JWT {
	SigningMethodHS512KMS = &SigningMethodHMAC{
		Name:      "HS512-KMS",
		Hash:      crypto.SHA512,
		KMSClient: kmsClient,
	}
	jwt.RegisterSigningMethod(SigningMethodHS512KMS.Alg(), func() jwt.SigningMethod {
		return SigningMethodHS512KMS
	})

	return &JWT{}
}

func (instance *JWT) Sign(claims JWTCustomClaims) (string, error) {
	token := jwt.NewWithClaims(SigningMethodHS512KMS, claims)

	ss, err := token.SignedString(nil)
	if err != nil {
		return "", WrapError(err)
	}

	return ss, nil
}

func (instance *JWT) Verify(tokenString string) (JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		return JWTCustomClaims{}, WrapError(err)
	}

	var claims *JWTCustomClaims
	var ok bool
	if claims, ok = token.Claims.(*JWTCustomClaims); !ok {
		return JWTCustomClaims{}, fmt.Errorf("invalid claims")
	}

	if !token.Valid {
		return JWTCustomClaims{}, fmt.Errorf("invalid token")
	}

	return *claims, nil
}
