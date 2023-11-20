package mw

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	error_response "practise-task-service/pkg/handler/error-response"
)

const invalidExtension string = "Неправильный формат файла"

func ExtensionValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			error_response.NewErrorResponse(w, r, http.StatusInternalServerError, "Не удалось обработать файл")
			return
		}
		if extension(fileHeader) != ".docx" || extension(fileHeader) != ".doc" {
			error_response.NewErrorResponse(w, r, http.StatusBadRequest, invalidExtension)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func extension(header *multipart.FileHeader) string {
	return filepath.Ext(header.Filename)
}
