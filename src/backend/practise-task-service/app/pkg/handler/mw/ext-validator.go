package mw

import (
	"github.com/go-chi/render"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func ExtensionValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, map[string]interface{}{
				"Ошибка ": err,
			})
			return
		}
		if extension(fileHeader) != ".docx" || extension(fileHeader) != ".doc" {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, map[string]interface{}{
				"Ошибка": "Неправильный формат файла",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func extension(header *multipart.FileHeader) string {
	return filepath.Ext(header.Filename)
}
