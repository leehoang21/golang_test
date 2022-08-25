package hashpassword

import (
	"golang.org/x/crypto/bcrypt"
)

type HashPassword interface {
	GererateHashedPassword() (string, error)
	ComparePassword(value string) error
	String() string
	Set(string)
}

type p string

func (p p) GererateHashedPassword() (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(hashed), err
}

func (p p) String() string {
	return string(p)
}

func (v p) Set(val string) {
	v = p(val)
}

func (pOld p) ComparePassword(pNew p) error {
	return bcrypt.CompareHashAndPassword([]byte(pOld), []byte(pNew))
}
