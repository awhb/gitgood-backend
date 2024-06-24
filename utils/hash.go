package utils

import "golang.org/x/crypto/bcrypt"

// Generates a hashed password using the bcrypt algorithm and returns 
// the hashed password as a string.
// If there is an error during the hashing process, it returns an error.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Returns a boolean indicating whether the password matches the hash.
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
