package mw

import (
	"log/slog"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"practise-task-service/pkg/error-response"
)

const invalidFileKey string = "Неправильный ключ multipart form"
const invalidExtension string = "Неправильный формат файла"

func ExtensionValidator(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("компонент", "mw/проверка расширения"))

		log.Info("middleware Проверка расширения включен")

		fn := func(w http.ResponseWriter, r *http.Request) {
			_, fileHeader, err := r.FormFile("file")
			if err != nil {
				log.Info("неправильный ключ файла multipart form")
				error_response.New(w, r, http.StatusBadRequest, invalidFileKey)
				return
			}

			ext := extension(fileHeader)

			if ext != ".docx" && ext != ".doc" {
				log.Info("неправильное расширение файла",
					slog.String("расширение", ext))
				error_response.New(w, r, http.StatusBadRequest, invalidExtension)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func extension(header *multipart.FileHeader) string {
	return filepath.Ext(header.Filename)
}
