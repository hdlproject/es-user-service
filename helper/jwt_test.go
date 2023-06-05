package helper

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestJWT_Sign(t *testing.T) {
	kmsClient, err := NewKMSClient("AKIAWZQRAROIO2J2V7VE", "CGPbr6mcUcIKskdnAo6uv1CuyXtH+iX6f5bapwCb")
	if err != nil {
		t.Fatal(err)
	}

	jwtGenerator := NewJWT(kmsClient)

	expectedClaims := JWTCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "es-user-service",
			Subject:   "es-services",
			Id:        uuid.New().String(),
			Audience:  "es-services",
		},
	}
	signedString, err := jwtGenerator.Sign(expectedClaims)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := jwtGenerator.Verify(signedString)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expectedClaims, claims); diff != "" {
		t.Fatalf("(-want/+got)\n%s", diff)
	}
}
