package hashpassword

import (
	"golang.org/x/crypto/bcrypt"
)

// type HashPassword interface {
// 	GenerateHashedPassword() (string, error)
// 	ComparePassword(value p) error
// 	String() string
// 	Set(string)
// }

type Password string

func NewPassword(pass string) Password {
	return Password(pass)
}
func (p Password) GenerateHashedPassword() (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(hashed), err
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Set(val string) {
	p = Password(val)
}

func (p Password) ComparePassword(pNew Password) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(pNew))
}
