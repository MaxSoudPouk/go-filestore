package repositories

import (
	"errors"
	"go-filestore/domain/entities"
	"go-filestore/domain/interfaces"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bucketRepo struct {
	db *gorm.DB
}

func (b bucketRepo) CreateBucket(bucket entities.Bucket) (*entities.Bucket, error) {
	check, err := b.CheckBucketExist(bucket.BucketName)
	if err != nil {
		return nil, err
	}
	if !check {
		subfolder := filepath.Join("./buckets", bucket.BucketName)
		if err := os.MkdirAll(subfolder, os.ModePerm); err != nil {
			return nil, err
		}
	}
	err = b.db.Where("bucket_name = ?", bucket.BucketName).Preload(clause.Associations).FirstOrCreate(&bucket).Error
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (b bucketRepo) RemoveBucket(bucket string) error {
	subfolder := filepath.Join("./buckets", bucket)
	if err := os.RemoveAll(subfolder); err != nil {
		return err
	}
	err := b.db.Where("bucket_name = ?", bucket).Delete(&entities.Bucket{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (b bucketRepo) GetBucketByName(bucketname string) (*entities.Bucket, error) {
	var bucket entities.Bucket
	err := b.db.Where("bucket_name = ?", bucketname).Preload(clause.Associations).First(&bucket).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("bucket dose not exist")
		}
		return nil, err
	}
	return &bucket, nil
}

func (b bucketRepo) CheckBucketExist(bucket string) (bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}

	// Construct the path to the "buckets" folder
	bucketsPath := filepath.Join(dir, "buckets")

	// Check if the "test" folder exists within the "buckets" folder
	testFolderPath := filepath.Join(bucketsPath, bucket)
	_, err = os.Stat(testFolderPath)
	if err != nil {
		if os.IsNotExist(err) {
			// "test" folder does not exist
			return false, nil
		}
		// Some other error occurred
		return false, err
	}

	// "test" folder exists
	return true, nil
}

func (b bucketRepo) GetBuckets() ([]entities.Item, error) {
	rootFolder := "./buckets/"
	var items []entities.Item

	err := filepath.Walk(rootFolder, func(path string, info fs.FileInfo, err error) error {
		// Skip the root folder
		if path == rootFolder {
			return nil
		}

		// Skip non-directories
		if !info.IsDir() {
			return nil
		}

		item := entities.Item{Name: info.Name(), Type: "directory"}
		items = append(items, item)

		return filepath.SkipDir // Skip traversing subdirectories
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (b bucketRepo) GetBucketItems(bucketname string) ([]entities.Item, error) {
	rootFolder := "./buckets/" + bucketname
	var items []entities.Item
	err := filepath.WalkDir(rootFolder, func(path string, d fs.DirEntry, err error) error {
		// Skip the root folder
		if path == rootFolder {
			return nil
		}

		item := entities.Item{Name: d.Name()}
		if d.IsDir() {
			item.Type = "directory"
		} else {
			item.Type = "file"
			item.FileType = strings.TrimPrefix(filepath.Ext(d.Name()), ".")
		}

		items = append(items, item)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func NewBucketRepo(db *gorm.DB) interfaces.BucketRepo {
	db.AutoMigrate(&entities.BucketKey{})
	db.AutoMigrate(&entities.Bucket{})
	return &bucketRepo{
		db: db,
	}
}
