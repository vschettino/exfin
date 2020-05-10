package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id       uint   `json:"id" pg:",pk"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (a Account) String() string {
	return fmt.Sprintf("Account<%d %s %s>", a.Id, a.Name, a.Email)
}

func (a *Account) VerifyPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pwd))
	return err == nil
}

func (a *Account) SetHashPassword(pwd string) {
	hash, _ := HashPassword(pwd)
	a.Password = hash
}

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}
