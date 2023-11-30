package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"practise-task-service/internal/models"
	"practise-task-service/internal/usecase"
	"practise-task-service/pkg/error-response"
	"strconv"
	"strings"
)

const DocxExt = ".docx"

func (h *Handler) GetAllPracticeTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		practices, err := h.service.PracticeGetter.GetGroup()
		if errors.Is(err, usecase.ErrNoResult) {
			error_response.New(w, r, http.StatusNoContent, err.Error())
			return
		}
		if err != nil {
			error_response.New(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		render.JSON(w, r, practices)
	}
}

func (h *Handler) GetPractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getId(r)
		if err != nil {
			error_response.New(w, r, http.StatusBadRequest, err.Error())
			return
		}

		practice, err := h.service.PracticeGetter.Get(id)
		if err != nil {
			error_response.New(w, r, http.StatusInternalServerError, "не удалось получить практическую")
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+practice.Theme+DocxExt)

		http.ServeFile(w, r, practice.File.Name())
		defer practice.File.Close()
	}
}

func (h *Handler) SearchPractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		academicSubject := r.URL.Query().Get("item")

		practicesInfo, err := h.service.PracticeGetter.GetBySearch(title, academicSubject)
		if errors.Is(err, usecase.ErrNoResult) {
			error_response.New(w, r, http.StatusNoContent, err.Error())
			return
		}
		if err != nil {
			error_response.New(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		render.JSON(w, r, practicesInfo)
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
			Author:          r.FormValue("author"),
			Theme:           r.FormValue("theme"),
			AcademicSubject: r.FormValue("academicSubject"),
			AccessGroup:     strings.Split(r.FormValue("accessGroup"), ","),
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			error_response.New(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		defer file.Close()

		upload.File = file
		upload.FileSize = handler.Size

		id, err := h.service.Save(upload, handler)
		if err != nil {
			error_response.New(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		render.JSON(w, r, map[string]interface{}{
			"ID сохранненого файла - ": id,
		})

	}
}

func (h *Handler) DeletePractice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getId(r)
		if err != nil {
			error_response.New(w, r, http.StatusBadRequest, err.Error())
			return
		}

		err = h.service.PracticeDeleter.Delete(id)
		if err != nil {
			error_response.New(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"ID удаленной практической:": id,
		})
	}
}

func getId(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		return 0, errors.New("empty id")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return id, nil
}
