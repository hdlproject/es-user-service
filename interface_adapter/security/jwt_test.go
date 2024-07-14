package security

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestJWT_Sign(t *testing.T) {
	kmsClient, err := NewKMSClient("asd", "asd")
	if err != nil {
		t.Fatal(err)
	}

	jwtGenerator := NewJWT(kmsClient)

	tests := []struct {
		name          string
		signingMethod jwt.SigningMethod
	}{
		{
			name:          "symmetric",
			signingMethod: signingMethodHS512KMS,
		},
		{
			name:          "asymmetric",
			signingMethod: signingMethodRS512KMS,
		},
		{
			name:          "asymmetric with offline verification",
			signingMethod: signingMethodRS512KMSOffline,
		},
	}

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			signedString, err := jwtGenerator.Sign(expectedClaims, test.signingMethod)
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
		})
	}
}
