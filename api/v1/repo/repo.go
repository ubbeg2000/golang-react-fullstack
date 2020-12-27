package repo

import (
	db "github.com/ubbeg2000/golang-react-fullstack/api/v1/database"
)

type Repo struct {
	User  UserRepo
	Event EventRepo
}

func New() Repo {
	if conn, err := db.New(); err == nil {
		return Repo{
			UserRepo{conn},
			EventRepo{conn},
		}
	} else {
		panic("Failed to connect to db")
	}
}
