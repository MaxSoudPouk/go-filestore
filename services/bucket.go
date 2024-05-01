package services

import (
	"errors"
	"go-filestore/domain/entities"
	"go-filestore/domain/interfaces"
	"go-filestore/domain/models"

	"github.com/google/uuid"
)

type bucketsvc struct {
	bucketrepo interfaces.BucketRepo
}

func (b bucketsvc) CreateBucket(bucket models.Bucket) (*models.Bucket, error) {
	bucketentity := entities.Bucket{
		BucketName: bucket.BucketName,
		BucketKey: entities.BucketKey{
			Key: uuid.New().String(),
		},
	}
	bucketrepo, err := b.bucketrepo.CreateBucket(bucketentity)
	if err != nil {
		return nil, err
	}
	result := models.Bucket{
		ID:         bucketrepo.ID,
		BucketName: bucketrepo.BucketName,
		BucketKey:  bucketrepo.BucketKey.Key,
	}
	return &result, nil
}

func (b bucketsvc) DeleteBucket(bucket string) error {
	check, err := b.bucketrepo.CheckBucketExist(bucket)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("bucket dose not exist")
	}
	err = b.bucketrepo.RemoveBucket(bucket)
	if err != nil {
		return err
	}
	return nil
}

func (b bucketsvc) BucketKeyVerify(bucketname string, key string) error {
	bucket, err := b.bucketrepo.GetBucketByName(bucketname)
	if err != nil {
		return err
	}
	if bucket.BucketKey.Key == "" || bucket.BucketKey.Key != key {
		return errors.New("key invalid")
	}
	return nil
}

func (b bucketsvc) GetBuckets() ([]entities.Item, error) {
	items, err := b.bucketrepo.GetBuckets()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (b bucketsvc) GetBucketItems(bucket string) ([]entities.Item, error) {
	items, err := b.bucketrepo.GetBucketItems(bucket)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (b bucketsvc) GetBucketByName(bucket string) (*models.Bucket, error) {
	bucketrepo, err := b.bucketrepo.GetBucketByName(bucket)
	if err != nil {
		return nil, err
	}
	result := models.Bucket{
		ID:         bucketrepo.ID,
		BucketName: bucketrepo.BucketName,
		BucketKey:  bucketrepo.BucketKey.Key,
	}
	return &result, nil
}

func NewBucketSvc(bucketrepo interfaces.BucketRepo) interfaces.BucketSvc {
	return &bucketsvc{
		bucketrepo: bucketrepo,
	}
}
