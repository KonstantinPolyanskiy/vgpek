package error_response

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResp struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.WriteHeader(code)
	render.JSON(w, r, ErrResp{Message: message})
}
