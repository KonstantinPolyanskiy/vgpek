package handler

import (
	"mod/internal/models"
	"net/http"
	"strings"
)

type PracticeSaver interface {
	Save(request models.UploadPracticeRequest) (int, error)
}

func (h *Handler) GetAllPracticeTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) GetPractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) SearchPractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) UploadPractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, "Слишком большой размер файла", http.StatusBadRequest)
		}

		upload := models.UploadPracticeRequest{
			Title:           r.FormValue("title"),
			Theme:           r.FormValue("theme"),
			AcademicSubject: r.FormValue("academicSubject"),
			AccessGroup:     strings.Split(r.FormValue("accessGroup"), ","),
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Не удалось получить файл!", http.StatusBadRequest)
		}
		defer file.Close()

		upload.File = file
		upload.FileSize = handler.Size
	}
}

func (h *Handler) DeletePractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
