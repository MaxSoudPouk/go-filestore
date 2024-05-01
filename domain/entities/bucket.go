package entities

import "gorm.io/gorm"

type Bucket struct {
	gorm.Model
	BucketName string
	BucketKey  BucketKey
}

type BucketKey struct {
	gorm.Model
	BucketID uint `gorm:"unique"`
	Key      string
}

type Item struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	FileType string `json:"file_type"`
}
