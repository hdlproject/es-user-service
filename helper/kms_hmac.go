package helper

import (
	"crypto"

	"github.com/golang-jwt/jwt"
)

type SigningMethodHMAC struct {
	Name      string
	Hash      crypto.Hash
	KMSClient *KMSClient
}

var (
	SigningMethodHS512KMS *SigningMethodHMAC
)

func (m *SigningMethodHMAC) Alg() string {
	return m.Name
}

func (m *SigningMethodHMAC) Verify(signingString, signature string, key interface{}) error {
	sig, err := jwt.DecodeSegment(signature)
	if err != nil {
		return err
	}

	if m.KMSClient == nil {
		return jwt.ErrHashUnavailable
	}

	err = m.KMSClient.VerifyMac(string(sig), signingString)
	if err != nil {
		return WrapError(err)
	}

	return nil
}

func (m *SigningMethodHMAC) Sign(signingString string, key interface{}) (string, error) {
	if m.KMSClient == nil {
		return "", jwt.ErrHashUnavailable
	}

	hashStr, err := m.KMSClient.GenerateMacInput(signingString)
	if err != nil {
		return "", WrapError(err)
	}

	return jwt.EncodeSegment([]byte(hashStr)), nil
}
