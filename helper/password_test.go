package helper

import (
	"testing"

	"github.com/google/uuid"
)

func TestPasswordManager_IsPasswordValid(t *testing.T) {
	passwordManager := NewPasswordManager(uuid.NewString())

	password := uuid.NewString()

	hashedPassword := passwordManager.GetHashedPassword(password)

	tests := []struct {
		name           string
		password       string
		hashedPassword string
		isValid        bool
	}{
		{
			name:           "valid",
			password:       password,
			hashedPassword: hashedPassword,
			isValid:        true,
		},
		{
			name:           "invalid",
			password:       uuid.NewString(),
			hashedPassword: hashedPassword,
			isValid:        false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isValid := passwordManager.IsPasswordValid(test.password, test.hashedPassword)
			if isValid != test.isValid {
				t.Fatalf("expect is_valid %v but got %v", test.isValid, isValid)
			}
		})
	}
}
