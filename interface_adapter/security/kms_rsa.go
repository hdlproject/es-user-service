package security

import (
	"crypto"
	"crypto/rsa"

	"github.com/golang-jwt/jwt"

	"github.com/hdlproject/es-user-service/helper"
)

type signingMethodRSA struct {
	name                  string
	kmsClient             *KMSClient
	useOnlineVerification bool
	hash                  crypto.Hash
	key                   interface{}
}

var (
	signingMethodRS512KMS        *signingMethodRSA
	signingMethodRS512KMSOffline *signingMethodRSA
)

func (m *signingMethodRSA) Alg() string {
	return m.name
}

func (m *signingMethodRSA) Verify(signingString, signature string, key interface{}) error {
	if m.useOnlineVerification {
		return m.verifyOnline(signingString, signature)
	}

	return m.verifyOffline(signingString, signature)
}

func (m *signingMethodRSA) verifyOnline(signingString, signature string) error {
	sig, err := jwt.DecodeSegment(signature)
	if err != nil {
		return err
	}

	if m.kmsClient == nil {
		return jwt.ErrHashUnavailable
	}

	err = m.kmsClient.Verify(string(sig), signingString)
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (m *signingMethodRSA) verifyOffline(signingString, signature string) error {
	var err error

	var sig []byte
	if sig, err = jwt.DecodeSegment(signature); err != nil {
		return err
	}

	var rsaKey *rsa.PublicKey
	var ok bool

	if rsaKey, ok = m.key.(*rsa.PublicKey); !ok {
		return jwt.ErrInvalidKeyType
	}

	if !m.hash.Available() {
		return jwt.ErrHashUnavailable
	}
	hasher := m.hash.New()
	hasher.Write([]byte(signingString))

	return rsa.VerifyPKCS1v15(rsaKey, m.hash, hasher.Sum(nil), sig)
}

func (m *signingMethodRSA) Sign(signingString string, key interface{}) (string, error) {
	if m.kmsClient == nil {
		return "", jwt.ErrHashUnavailable
	}

	hashStr, err := m.kmsClient.Sign(signingString)
	if err != nil {
		return "", helper.WrapError(err)
	}

	return jwt.EncodeSegment([]byte(hashStr)), nil
}
