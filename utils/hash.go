package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a given plain text password using bcrypt.
// It returns the hashed password as a string and any error encountered during the hashing process.
// The bcrypt cost factor is set to 14 for a good balance between security and performance.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CompareHashAndPassword compares a hashed password with its possible plain text equivalent.
// It returns true if the passwords match and false otherwise. Any error during the comparison
// process will result in a return value of false.
func CompareHashAndPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil // return false if error else return true
}
