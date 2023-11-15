package models

import "mime/multipart"

type UploadPracticeRequest struct {
	Title           string   `json:"title"`
	Theme           string   `json:"theme"`
	AcademicSubject string   `json:"academicSubject"`
	AccessGroup     []string `json:"accessGroup"`
	File            multipart.File
	FileSize        int64
}
