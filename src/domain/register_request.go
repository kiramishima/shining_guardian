package domain

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

type RegisterRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (u *RegisterRequest) Hash256Password(password string) string {
	buf := []byte(password)
	pwd := sha3.New256()
	pwd.Write(buf)
	return hex.EncodeToString(pwd.Sum(nil))
}
