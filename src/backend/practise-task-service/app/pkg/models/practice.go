package models

import "os"

type PracticeResponse struct {
	PracticeInfo
	PracticeFile
}

type PracticeInfo struct {
	Author          string `json:"author" db:"author"`
	Title           string `json:"title" db:"title"`
	Theme           string `json:"theme" db:"theme"`
	AcademicSubject string `json:"academicSubject" db:"academic_subject"`
}

type PracticeFile struct {
	File os.File
}
