package models

type Bucket struct {
	ID         uint   `json:"bucket_id"`
	BucketName string `json:"bucket_name"`
	BucketKey  string `json:"bucket_key"`
}
