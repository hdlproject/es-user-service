package security

import (
	"crypto"
	"crypto/rsa"

	"github.com/golang-jwt/jwt"

	"github.com/hdlproject/es-user-service/helper"
)

type SigningMethodRSA struct {
	Name                  string
	KMSClient             *KMSClient
	UseOnlineVerification bool
	Hash                  crypto.Hash
	Key                   interface{}
}

var (
	SigningMethodRS512KMS        *SigningMethodRSA
	SigningMethodRS512KMSOffline *SigningMethodRSA
)

func (m *SigningMethodRSA) Alg() string {
	return m.Name
}

func (m *SigningMethodRSA) Verify(signingString, signature string, key interface{}) error {
	if m.UseOnlineVerification {
		return m.verifyOnline(signingString, signature)
	}

	return m.verifyOffline(signingString, signature)
}

func (m *SigningMethodRSA) verifyOnline(signingString, signature string) error {
	sig, err := jwt.DecodeSegment(signature)
	if err != nil {
		return err
	}

	if m.KMSClient == nil {
		return jwt.ErrHashUnavailable
	}

	err = m.KMSClient.Verify(string(sig), signingString)
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (m *SigningMethodRSA) verifyOffline(signingString, signature string) error {
	var err error

	var sig []byte
	if sig, err = jwt.DecodeSegment(signature); err != nil {
		return err
	}

	var rsaKey *rsa.PublicKey
	var ok bool

	if rsaKey, ok = m.Key.(*rsa.PublicKey); !ok {
		return jwt.ErrInvalidKeyType
	}

	if !m.Hash.Available() {
		return jwt.ErrHashUnavailable
	}
	hasher := m.Hash.New()
	hasher.Write([]byte(signingString))

	return rsa.VerifyPKCS1v15(rsaKey, m.Hash, hasher.Sum(nil), sig)
}

func (m *SigningMethodRSA) Sign(signingString string, key interface{}) (string, error) {
	if m.KMSClient == nil {
		return "", jwt.ErrHashUnavailable
	}

	hashStr, err := m.KMSClient.Sign(signingString)
	if err != nil {
		return "", helper.WrapError(err)
	}

	return jwt.EncodeSegment([]byte(hashStr)), nil
}
