package route

import (
	"go-filestore/api/controllers"
	"go-filestore/repositories"
	"go-filestore/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFileRoute(router fiber.Router, db *gorm.DB) {
	fr := repositories.NewFileRepo()
	br := repositories.NewBucketRepo(db)
	fs := services.NewFileSvc(fr, br)
	bs := services.NewBucketSvc(br)
	fc := controllers.NewFileCtrl(fs, bs)

	api := router.Group("file")

	api.Post("/upload", fc.Uploadfile)
	api.Delete("/delete", fc.Deletefile)
}
