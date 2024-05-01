package services

import (
	"errors"
	"go-filestore/api/middleware"
	"go-filestore/domain/entities"
	"go-filestore/domain/interfaces"
	"go-filestore/domain/models"
	"go-filestore/utilities"
)

type usersvc struct {
	userrepo interfaces.UserRepo
}

func (u usersvc) Login(user models.User) (*middleware.TokenPair, error) {
	userdata, err := u.userrepo.GetUser(user.Username)
	if err != nil {
		return nil, err
	}
	password, err := utilities.GetAESEncrypted(user.Password)
	if err != nil {
		return nil, err
	}
	if userdata.Password == password {
		TokenPair, err := middleware.GenerateJWTToken(user.Username)
		if err != nil {
			return nil, err
		}
		return TokenPair, nil

	}
	return nil, errors.New("invalid credential")
}

func (u usersvc) CreateUser(user models.User) (*models.User, error) {
	password, err := utilities.GetAESEncrypted(user.Password)
	if err != nil {
		return nil, err
	}
	userentity := entities.User{
		Username: user.Username,
		Password: password,
	}
	userdata, err := u.userrepo.CreateUser(userentity)
	if err != nil {
		return nil, err
	}
	result := models.User{
		Username: userdata.Username,
	}
	return &result, nil
}

func (u usersvc) UserCheck() (*bool, error) {
	check, err := u.userrepo.CheckUserExisted()
	if err != nil {
		return nil, err
	}
	return check, nil
}

func NewUserSvc(userrepo interfaces.UserRepo) interfaces.UserSvc {
	return &usersvc{
		userrepo: userrepo,
	}
}
