package apis

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/sant470/devetron/apis/v1"
	"github.com/sant470/devetron/utils/errors"
)

type Handler func(rw http.ResponseWriter, r *http.Request) *errors.AppError

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := h(rw, r); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Message))
	}
}

func InitSerachRoutes(r *chi.Mux, sh *v1.SearchHandler) {
	r.Route("/api/v1/search", func(r chi.Router) {
		r.Method(http.MethodPost, "/", Handler(sh.Search))
	})
}
