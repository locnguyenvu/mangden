package appconfig

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locnguyenvu/mangden/pkg/httpresponse"
	"github.com/locnguyenvu/mangden/proto"
)

type handler struct {
	acRepository Repository
}

func NewHandler(repository Repository) http.Handler {
	r := chi.NewRouter()

	h := &handler{acRepository: repository}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		appconfig := repository.GetByName("name")
		resp := &proto.Success{
			Status:  appconfig.Name,
			Message: appconfig.Value,
		}
		httpresponse.SendSuccess(w, r, resp)
	})

	r.Get("/{configID}", h.detailHandler)

	return r
}

func (h handler) detailHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "configID")

	var result AppConfig

	h.acRepository.DB().First(&result, id)

	httpresponse.SendSuccess(w, r, &proto.ConfigDetail{
		Name:      result.Name,
		Value:     result.Value,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	})
}
