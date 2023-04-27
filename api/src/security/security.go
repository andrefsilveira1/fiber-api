package security

import "golang.org/x/crypto/bcrypt"

// Receive the hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

// Make a verification of password with the hash
func Verify(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))

}
