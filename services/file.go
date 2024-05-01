package services

import (
	"errors"
	"go-filestore/domain/interfaces"
	"mime/multipart"
)

type fileSvc struct {
	filerepo   interfaces.FileRepo
	bucketrepo interfaces.BucketRepo
}

func (f fileSvc) Uploadfile(bucket string, file multipart.FileHeader) (*string, error) {
	check, err := f.bucketrepo.CheckBucketExist(bucket)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, errors.New("bucket dose not exist")
	}
	filename, err := f.filerepo.Uploadfile(bucket, file)
	if err != nil {
		return nil, err
	}
	return filename, nil
}

func (f fileSvc) Deletefile(bucket string, filename string) error {
	check, err := f.bucketrepo.CheckBucketExist(bucket)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("bucket dose not exist")
	}
	err = f.filerepo.Deletefile(bucket, filename)
	if err != nil {
		return err
	}
	return nil
}

func NewFileSvc(filerepo interfaces.FileRepo, bucketrepo interfaces.BucketRepo) interfaces.FileSvc {
	return &fileSvc{
		filerepo:   filerepo,
		bucketrepo: bucketrepo,
	}
}
