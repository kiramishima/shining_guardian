package domain

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

type AuthRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

func (u *AuthRequest) Hash256Password(password string) string {
	buf := []byte(password)
	pwd := sha3.New256()
	pwd.Write(buf)
	return hex.EncodeToString(pwd.Sum(nil))
}
