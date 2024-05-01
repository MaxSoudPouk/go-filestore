package repositories

import (
	"errors"
	"go-filestore/domain/entities"
	"go-filestore/domain/interfaces"

	"gorm.io/gorm"
)

type userrepo struct {
	db *gorm.DB
}

func (u userrepo) GetUser(user string) (*entities.User, error) {
	var userdata entities.User
	err := u.db.Where("username = ?", user).First(&userdata).Error
	if err != nil {
		return nil, err
	}
	return &userdata, nil
}

func (u userrepo) GetPassword(password string) (*string, error) {
	var password_res string
	err := u.db.Raw("select password(?) as password;", password).Scan(&password_res).Error
	if err != nil {
		return nil, err
	}
	return &password_res, nil
}

func (u userrepo) CreateUser(user entities.User) (*entities.User, error) {
	check, err := u.CheckUserExisted()
	if err != nil {
		return nil, err
	}
	if *check {
		return nil, errors.New("the system can have only one root user")
	}
	err = u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userrepo) CheckUserExisted() (*bool, error) {
	var usercnt int64
	var status = false
	err := u.db.Model(&entities.User{}).Count(&usercnt).Error
	if err != nil {
		return nil, err
	}
	if usercnt > 0 {
		status = true
	}
	return &status, nil
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	db.AutoMigrate(&entities.User{})
	return &userrepo{
		db: db,
	}
}
