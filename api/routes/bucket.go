package route

import (
	"go-filestore/api/controllers"
	"go-filestore/api/middleware"
	"go-filestore/repositories"
	"go-filestore/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewBuckerRoute(router fiber.Router, db *gorm.DB) {
	br := repositories.NewBucketRepo(db)
	bs := services.NewBucketSvc(br)
	bc := controllers.NewBucketCtrl(bs)

	api := router.Group("bucket", middleware.AccessToken)

	api.Post("/", bc.CreateBucket)
	api.Delete("/:bucket", bc.DeleteBucket)
	api.Get("/info/:bucket", bc.GetBucketByName)
	api.Get("/", bc.GetBuckets)
	api.Get("/get-items/:bucket", bc.GetBucketItems)
}
