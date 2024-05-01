package controllers

import (
	"go-filestore/api/middleware"
	"go-filestore/domain/interfaces"
	"go-filestore/domain/models"

	"github.com/gofiber/fiber/v2"
)

type userctrl struct {
	usersvc interfaces.UserSvc
}

func (u userctrl) Login(ctx *fiber.Ctx) error {
	body := models.User{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	token, err := u.usersvc.Login(body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, token)
}

func (u userctrl) CheckUser(ctx *fiber.Ctx) error {
	check, err := u.usersvc.UserCheck()
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, check)
}

func (u userctrl) CreateUser(ctx *fiber.Ctx) error {
	body := models.User{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	user, err := u.usersvc.CreateUser(body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, user)
}

func (u userctrl) RefreshToken(ctx *fiber.Ctx) error {
	token, err := middleware.GenerateRefreshToken(ctx)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessResponse(ctx, token)
}

func NewUserCtrl(usersvc interfaces.UserSvc) interfaces.UserCtrl {
	return &userctrl{
		usersvc: usersvc,
	}
}
