package helper

import "golang.org/x/crypto/argon2"

type PasswordManager struct {
	salt string
}

func NewPasswordManager(salt string) *PasswordManager {
	return &PasswordManager{
		salt: salt,
	}
}

func (instance *PasswordManager) GetHashedPassword(password string) string {
	hash := argon2.IDKey(
		[]byte(password),
		[]byte(instance.salt),
		1,
		64*1024,
		4,
		32,
	)

	return string(hash)
}

func (instance *PasswordManager) IsPasswordValid(password, hashedPassword string) bool {
	return instance.GetHashedPassword(password) == hashedPassword
}
