package view

import (
	"net/http"

	"github.com/locnguyenvu/mangden/pkg/view/template"
)

type View struct {
	r      *http.Request
	tpl    *template.Engine
	params map[string]interface{}
}

func New(r *http.Request, tpl *template.Engine) *View {
	params := make(map[string]interface{}, 0)
	return &View{r, tpl, params}
}

func (v *View) Render(viewName string) []byte {
	return v.tpl.Render(viewName+".html", v.params)
}

func (v *View) Set(param string, value interface{}) *View {
	v.params[param] = value
	return v
}
