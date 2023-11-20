package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"mime/multipart"
	"net/http"
	"practise-task-service/pkg/handler/error-response"
	"practise-task-service/pkg/models"
	"strconv"
	"strings"
)

const DocxExt = ".docx"

type PracticeSaver interface {
	Save(request models.UploadPracticeRequest, fh *multipart.FileHeader) (int, error)
}
type PracticeGetter interface {
	Get(id int) (models.PracticeFile, error)
}

func (h *Handler) GetAllPracticeTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) GetPractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			error_response.NewErrorResponse(w, r, http.StatusBadRequest, "пустой id в запросе")
			return
		}
		id, err := strconv.Atoi(idParam)
		if err != nil {
			error_response.NewErrorResponse(w, r, http.StatusBadRequest, "не удалось конвертовать в int")
			return
		}

		practice, err := h.service.PracticeGetter.Get(id)
		if err != nil {
			error_response.NewErrorResponse(w, r, http.StatusInternalServerError, "не удалось получить практическую")
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+practice.Theme+DocxExt)

		http.ServeFile(w, r, practice.File.Name())
		defer practice.File.Close()
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

		id, err := h.service.Save(upload, handler)
		if err != nil {
			render.JSON(w, r, map[string]interface{}{
				"ошибка - ": err,
			})
		}
		render.JSON(w, r, map[string]interface{}{
			"ID сохранненого файла - ": id,
		})

	}
}

func (h *Handler) DeletePractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
