package security

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"

	"github.com/hdlproject/es-user-service/helper"
)

type JWT struct {
}

type JWTCustomClaims struct {
	jwt.StandardClaims
}

func NewJWT(kmsClient *KMSClient) *JWT {
	SigningMethodHS512KMS = &SigningMethodHMAC{
		Name:      "HS512-KMS",
		KMSClient: kmsClient,
	}
	jwt.RegisterSigningMethod(SigningMethodHS512KMS.Alg(), func() jwt.SigningMethod {
		return SigningMethodHS512KMS
	})

	SigningMethodRS512KMS = &SigningMethodRSA{
		Name:                  "RS512-KMS",
		KMSClient:             kmsClient,
		UseOnlineVerification: true,
	}
	jwt.RegisterSigningMethod(SigningMethodRS512KMS.Alg(), func() jwt.SigningMethod {
		return SigningMethodRS512KMS
	})

	publicKeyStr, err := os.ReadFile("../../kms-public-key.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(publicKeyStr)
	b := block.Bytes
	key, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		panic(err)
	}

	SigningMethodRS512KMSOffline = &SigningMethodRSA{
		Name:      "RS512-KMS-Offline",
		KMSClient: kmsClient,
		Hash:      crypto.SHA512,
		Key:       key,
	}
	jwt.RegisterSigningMethod(SigningMethodRS512KMSOffline.Alg(), func() jwt.SigningMethod {
		return SigningMethodRS512KMSOffline
	})

	return &JWT{}
}

func (instance *JWT) Sign(claims JWTCustomClaims, method jwt.SigningMethod) (string, error) {
	token := jwt.NewWithClaims(method, claims)

	ss, err := token.SignedString(nil)
	if err != nil {
		return "", helper.WrapError(err)
	}

	return ss, nil
}

func (instance *JWT) Verify(tokenString string) (JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		return JWTCustomClaims{}, helper.WrapError(err)
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
