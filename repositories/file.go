package repositories

import (
	"errors"
	"go-filestore/domain/interfaces"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type fileRepos struct {
}

func (f fileRepos) Uploadfile(bucket string, file multipart.FileHeader) (*string, error) {
	path := fmt.Sprintf("./buckets/%s/%s", bucket, file.Filename)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			src, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer src.Close()
			// Create a new file on the server
			dst, err := os.Create(path)
			if err != nil {
				return nil, err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return nil, err
			}

			fmt.Printf("File uploaded: %s\n", file.Filename)
			return &file.Filename, nil
		}
	}
	return nil, errors.New("file with this name already exist")
}

func (f fileRepos) Deletefile(bucket string, filename string) error {
	path := fmt.Sprintf("./buckets/%s/%s", bucket, filename)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return errors.New("file dose not exist")
		}
		return err
	}
	// Delete the file
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func NewFileRepo() interfaces.FileRepo {
	return &fileRepos{}
}
