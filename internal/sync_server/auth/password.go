package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password string, passwordHash []byte) bool {
	return bcrypt.CompareHashAndPassword(passwordHash, []byte(password)) == nil
}
