package interfaces

import (
	"go-filestore/domain/entities"
	"go-filestore/domain/models"

	"github.com/gofiber/fiber/v2"
)

type BucketRepo interface {
	CreateBucket(bucket entities.Bucket) (*entities.Bucket, error)
	RemoveBucket(bucket string) error
	CheckBucketExist(bucket string) (bool, error)
	GetBucketByName(bucketname string) (*entities.Bucket, error)
	GetBuckets() ([]entities.Item, error)
	GetBucketItems(bucketname string) ([]entities.Item, error)
}

type BucketSvc interface {
	CreateBucket(bucket models.Bucket) (*models.Bucket, error)
	DeleteBucket(bucket string) error
	BucketKeyVerify(bucketname string, key string) error
	GetBuckets() ([]entities.Item, error)
	GetBucketItems(bucket string) ([]entities.Item, error)
	GetBucketByName(bucket string) (*models.Bucket, error)
}

type BucketCtrl interface {
	CreateBucket(ctx *fiber.Ctx) error
	DeleteBucket(ctx *fiber.Ctx) error
	GetBuckets(ctx *fiber.Ctx) error
	GetBucketItems(ctx *fiber.Ctx) error
	GetBucketByName(ctx *fiber.Ctx) error
}
