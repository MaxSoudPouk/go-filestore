package controllers

import (
	"go-filestore/api/middleware"
	"go-filestore/domain/interfaces"
	"go-filestore/domain/models"

	"github.com/gofiber/fiber/v2"
)

type bucketctrl struct {
	bucketsvc interfaces.BucketSvc
}

func (b bucketctrl) CreateBucket(ctx *fiber.Ctx) error {
	body := models.Bucket{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	if body.BucketName == "" {
		return middleware.NewErrorMessageResponse(ctx, "bucket name is required")
	}
	bucket, err := b.bucketsvc.CreateBucket(body)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, bucket)
}

func (b bucketctrl) DeleteBucket(ctx *fiber.Ctx) error {
	bucketname := ctx.Params("bucket")
	if bucketname == "" {
		return middleware.NewErrorMessageResponse(ctx, "bucket name is required")
	}
	err := b.bucketsvc.DeleteBucket(bucketname)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, "bucket deleted")
}

func (b bucketctrl) GetBuckets(ctx *fiber.Ctx) error {
	buckets, err := b.bucketsvc.GetBuckets()
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, buckets)
}

func (b bucketctrl) GetBucketItems(ctx *fiber.Ctx) error {
	bucket := ctx.Params("bucket")
	item, err := b.bucketsvc.GetBucketItems(bucket)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, item)
}

func (b bucketctrl) GetBucketByName(ctx *fiber.Ctx) error {
	bucket := ctx.Params("bucket")
	item, err := b.bucketsvc.GetBucketByName(bucket)
	if err != nil {
		return middleware.NewErrorResponses(ctx, err)
	}
	return middleware.NewSuccessMessageResponse(ctx, item)
}

func NewBucketCtrl(bucketsvc interfaces.BucketSvc) interfaces.BucketCtrl {
	return &bucketctrl{
		bucketsvc: bucketsvc,
	}
}
