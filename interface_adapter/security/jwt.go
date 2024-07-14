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
	signingMethodHS512KMS = &signingMethodHMAC{
		name:      "HS512-KMS",
		kmsClient: kmsClient,
	}
	jwt.RegisterSigningMethod(signingMethodHS512KMS.Alg(), func() jwt.SigningMethod {
		return signingMethodHS512KMS
	})

	signingMethodRS512KMS = &signingMethodRSA{
		name:                  "RS512-KMS",
		kmsClient:             kmsClient,
		useOnlineVerification: true,
	}
	jwt.RegisterSigningMethod(signingMethodRS512KMS.Alg(), func() jwt.SigningMethod {
		return signingMethodRS512KMS
	})

	publicKeyStr, err := os.ReadFile("../../kms-public-key.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(publicKeyStr)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	signingMethodRS512KMSOffline = &signingMethodRSA{
		name:      "RS512-KMS-Offline",
		kmsClient: kmsClient,
		hash:      crypto.SHA512,
		key:       key,
	}
	jwt.RegisterSigningMethod(signingMethodRS512KMSOffline.Alg(), func() jwt.SigningMethod {
		return signingMethodRS512KMSOffline
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
