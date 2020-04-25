package models

import "fmt"

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *Account
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}
