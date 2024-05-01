package controllers

import (
	"go-filestore/api/middleware"
	"go-filestore/domain/interfaces"
	"go-filestore/domain/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type fileCtrl struct {
	filesvc   interfaces.FileSvc
	bucketsvc interfaces.BucketSvc
}

func (f fileCtrl) Uploadfile(ctx *fiber.Ctx) error {
	var result models.File
	body := models.File{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if body.Bucket == "" {
		return middleware.NewErrorMessageResponse(ctx, "bucket name is required")
	}
	key := ctx.Get("x-api-key")
	println("key ", key)
	err = f.bucketsvc.BucketKeyVerify(body.Bucket, key)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}

	// Get the files from form
	files := form.File["FILES"]
	if len(files) == 0 {
		return middleware.NewErrorMessageResponse(ctx, "no file input")
	}
	for _, file := range files {
		filename, err := f.filesvc.Uploadfile(body.Bucket, *file)
		if err != nil {
			return middleware.NewErrorResponses(ctx, err)
		}
		result.Bucket = body.Bucket
		result.FileName = *filename
		result.FilePath = fmt.Sprintf("/buckets/%s/%s", body.Bucket, *filename)
	}
	return middleware.NewSuccessResponse(ctx, result)
}

func (f fileCtrl) Deletefile(ctx *fiber.Ctx) error {
	body := models.File{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	key := ctx.Get("x-api-key")
	println("key ", key)
	err = f.bucketsvc.BucketKeyVerify(body.Bucket, key)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	err = f.filesvc.Deletefile(body.Bucket, body.FileName)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, "file deleted")
}

func NewFileCtrl(filesvc interfaces.FileSvc, bucketsvc interfaces.BucketSvc) interfaces.FileCtrl {
	return &fileCtrl{
		filesvc:   filesvc,
		bucketsvc: bucketsvc,
	}
}
