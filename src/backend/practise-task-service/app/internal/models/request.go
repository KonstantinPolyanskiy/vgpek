package models

import "mime/multipart"

type UploadPracticeRequest struct {
	Title           string   `json:"title"`
	Author          string   `json:"author"`
	Theme           string   `json:"theme"`
	AcademicSubject string   `json:"academicSubject"`
	AccessGroup     []string `json:"accessGroup"`
	File            multipart.File
	FileSize        int64
}
