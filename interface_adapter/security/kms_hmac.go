package security

import (
	"github.com/golang-jwt/jwt"

	"github.com/hdlproject/es-user-service/helper"
)

type SigningMethodHMAC struct {
	Name      string
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
		return helper.WrapError(err)
	}

	return nil
}

func (m *SigningMethodHMAC) Sign(signingString string, key interface{}) (string, error) {
	if m.KMSClient == nil {
		return "", jwt.ErrHashUnavailable
	}

	hashStr, err := m.KMSClient.GenerateMac(signingString)
	if err != nil {
		return "", helper.WrapError(err)
	}

	return jwt.EncodeSegment([]byte(hashStr)), nil
}
