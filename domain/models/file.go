package models

type File struct {
	Bucket   string `json:"bucket"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
