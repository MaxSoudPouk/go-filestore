package route

import (
	"go-filestore/api/controllers"
	"go-filestore/api/middleware"
	"go-filestore/repositories"
	"go-filestore/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewUserRoute(router fiber.Router, db *gorm.DB) {
	ur := repositories.NewUserRepo(db)
	us := services.NewUserSvc(ur)
	uc := controllers.NewUserCtrl(us)

	api := router.Group("/user")

	api.Post("/", uc.CreateUser)
	api.Get("/check", uc.CheckUser)
	api.Post("/login", uc.Login)
	api.Post("/refresh", middleware.AccessRefreshToken, uc.RefreshToken)
}
