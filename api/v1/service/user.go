package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/auth"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/types"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/repo"
)

type UserService struct {
	repo repo.Repo
}

// FindAll : find all users that fulfills set parameters
func (s *UserService) FindAll(q *types.UserQuery) []models.User {
	return s.repo.User.FindMany(q.Limit, q.Page)
}

// FindByID : find a user by the user's ID
func (s *UserService) FindByID(id int) models.User {
	var user models.User
	user.ID = uint64(id)
	return s.repo.User.FindOne(user)
}

// FindByUsernameOrEmail : find a user by the user's username or email
func (s *UserService) FindByUsernameOrEmail(username string) models.User {
	if u := s.repo.User.FindOne(models.User{Username: username}); u.ID != 0 {
		return u
	}

	if e := s.repo.User.FindOne(models.User{Username: username}); e.ID != 0 {
		return e
	}

	return models.User{}
}

// Register : register a new user
func (s *UserService) Register(param models.User) (models.User, error) {
	if hash, err := auth.Encrypt(param.Password); err == nil {
		param.Password = hash
		return s.repo.User.Create(param)
	} else {
		return param, err
	}
}

// Update : update a user's data (ID is mandatory)
func (s *UserService) Update(param models.User) (models.User, error) {
	if param.Password != "" {
		if hash, err := auth.Encrypt(param.Password); err == nil {
			param.Password = hash
		} else {
			return param, err
		}
	}

	if res, err := s.repo.User.Update(param); err == nil {
		return res, nil
	} else {
		return param, err
	}
}

// Delete : delete a user by their id
func (s *UserService) Delete(param models.User) error {
	return s.repo.User.Delete(param)
}
