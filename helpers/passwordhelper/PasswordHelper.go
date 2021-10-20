package passwordhelper

import "golang.org/x/crypto/bcrypt"

func PasswordBcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func PasswordBcryptVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
