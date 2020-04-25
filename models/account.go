package models

import "fmt"

type Account struct  {
	Id       uint `json:"id" pg:",pk"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string
}

func (u Account) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Name, u.Email)
}
