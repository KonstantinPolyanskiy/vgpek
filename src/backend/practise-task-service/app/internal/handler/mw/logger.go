package mw

import (
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"time"
)

func Logger(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("компонент", "mw/Логгер"))

		log.Info("middleware Логгер включен")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("http метод", r.Method),
				slog.String("путь", r.URL.Path),
				slog.String("user agent", r.UserAgent()),
				slog.String("id реквеста", middleware.GetReqID(r.Context())))

			rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Info("запрос обработан",
					slog.Int("статус", rw.Status()),
					slog.String("время обработки", time.Since(t1).String()))
			}()

			next.ServeHTTP(rw, r)
		}
		return http.HandlerFunc(fn)
	}
}
