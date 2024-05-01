package interfaces

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type FileRepo interface {
	Uploadfile(bucket string, file multipart.FileHeader) (*string, error)
	Deletefile(bucket string, filename string) error
}

type FileSvc interface {
	Uploadfile(bucket string, file multipart.FileHeader) (*string, error)
	Deletefile(bucket string, filename string) error
}

type FileCtrl interface {
	Uploadfile(ctx *fiber.Ctx) error
	Deletefile(ctx *fiber.Ctx) error
}
