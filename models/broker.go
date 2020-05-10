package models

import (
	"fmt"
)

type Broker struct {
	Id          uint   `json:"id" pg:",pk"`
	Name        string `json:"name"`
	Institution string `json:"institution"`
	AccountId   uint   `json:"account_id"`
	Account     *Account
}

func (a Broker) String() string {
	return fmt.Sprintf("Broker<%d %s %s %s: %d>", a.Id, a.Name, a.Institution, a.Account, a.AccountId)
}
