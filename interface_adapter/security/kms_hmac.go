package security

import (
	"github.com/golang-jwt/jwt"

	"github.com/hdlproject/es-user-service/helper"
)

type signingMethodHMAC struct {
	name      string
	kmsClient *KMSClient
}

var (
	signingMethodHS512KMS *signingMethodHMAC
)

func (m *signingMethodHMAC) Alg() string {
	return m.name
}

func (m *signingMethodHMAC) Verify(signingString, signature string, key interface{}) error {
	sig, err := jwt.DecodeSegment(signature)
	if err != nil {
		return err
	}

	if m.kmsClient == nil {
		return jwt.ErrHashUnavailable
	}

	err = m.kmsClient.VerifyMac(string(sig), signingString)
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (m *signingMethodHMAC) Sign(signingString string, key interface{}) (string, error) {
	if m.kmsClient == nil {
		return "", jwt.ErrHashUnavailable
	}

	hashStr, err := m.kmsClient.GenerateMac(signingString)
	if err != nil {
		return "", helper.WrapError(err)
	}

	return jwt.EncodeSegment([]byte(hashStr)), nil
}
