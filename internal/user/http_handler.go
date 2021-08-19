package user

import (
	"net/http"

	"github.com/locnguyenvu/mangden/pkg/view"
	"github.com/locnguyenvu/mangden/pkg/view/template"
)

type HttpHandler struct {
	userRepository Repository
	templateEngine *template.Engine
}

func NewHttpHandler(userRepository Repository, templateEngine *template.Engine) *HttpHandler {
	return &HttpHandler{userRepository, templateEngine}
}

func (h HttpHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	view := view.New(r, h.templateEngine)
	w.Write(view.Render("login"))
}
