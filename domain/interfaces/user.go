package interfaces

import (
	"go-filestore/api/middleware"
	"go-filestore/domain/entities"
	"go-filestore/domain/models"

	"github.com/gofiber/fiber/v2"
)

type UserRepo interface {
	GetUser(user string) (*entities.User, error)
	GetPassword(password string) (*string, error)
	CreateUser(user entities.User) (*entities.User, error)
	CheckUserExisted() (*bool, error)
}

type UserSvc interface {
	Login(user models.User) (*middleware.TokenPair, error)
	CreateUser(user models.User) (*models.User, error)
	UserCheck() (*bool, error)
}

type UserCtrl interface {
	Login(ctx *fiber.Ctx) error
	CheckUser(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}
