package models

type Comment struct {
	Base
	AuthorID uint64 `json:"-"`
	Author   User   `json:"author"`
	Body     string `json:"body"`
}
