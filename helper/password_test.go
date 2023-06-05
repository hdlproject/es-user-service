package helper

import (
	"testing"

	"github.com/google/uuid"
)

func TestPasswordManager_IsPasswordValid(t *testing.T) {
	tests := struct {
		name string
		isValid
	}{}

	passwordManager := NewPasswordManager(uuid.NewString())

	password := uuid.NewString()

	hashedPassword := passwordManager.GetHashedPassword(password)

	isValid := passwordManager.IsPasswordValid(password, hashedPassword)
	if !isValid {
		t.Fatalf("invalid password")
	}
}
