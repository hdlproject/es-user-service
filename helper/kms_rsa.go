package helper

import (
	"github.com/golang-jwt/jwt"
)

type SigningMethodRSA struct {
	Name      string
	KMSClient *KMSClient
}

var (
	SigningMethodRS512KMS *SigningMethodRSA
)

func (m *SigningMethodRSA) Alg() string {
	return m.Name
}

func (m *SigningMethodRSA) Verify(signingString, signature string, key interface{}) error {
	sig, err := jwt.DecodeSegment(signature)
	if err != nil {
		return err
	}

	if m.KMSClient == nil {
		return jwt.ErrHashUnavailable
	}

	err = m.KMSClient.Verify(string(sig), signingString)
	if err != nil {
		return WrapError(err)
	}

	return nil
}

func (m *SigningMethodRSA) Sign(signingString string, key interface{}) (string, error) {
	if m.KMSClient == nil {
		return "", jwt.ErrHashUnavailable
	}

	hashStr, err := m.KMSClient.Sign(signingString)
	if err != nil {
		return "", WrapError(err)
	}

	return jwt.EncodeSegment([]byte(hashStr)), nil
}
