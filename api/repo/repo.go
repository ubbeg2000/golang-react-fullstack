package repo

import (
	db "golang-api/api/database"
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
