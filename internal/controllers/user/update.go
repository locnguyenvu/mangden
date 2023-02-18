package user

import (
    "errors"
    "net/http"
    "strconv"

	"github.com/go-chi/chi/v5"
)

type UpdateUserRequestBody struct {
    Firstname string `json:"firstName" validate:"required"`
    Lastname string `json:"lastName" validate:"required"`
    Yob int `json:"yob" validate:"required"`
}


func (c *Controller) HandleUpdate() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64);
        if err != nil {
            c.r.RenderJSON(w, 400, errors.New("Invalid id"))
            return
        }
        payload := UpdateUserRequestBody{}
        if err := loadRequestBody(r, &payload); err != nil {
            c.r.RenderJSON(w, 400, errors.New("Invalid request"))
            return
        }
        u := c.userRepository.Get(id)
        if u == nil {
            c.r.RenderJSON(w, 404, errors.New("User not found"))
            return
        }
        u.FirstName = payload.Firstname
        u.LastName = payload.Lastname
        u.Yob = payload.Yob
        if err := c.userRepository.Save(u); err != nil {
            c.r.RenderJSON(w, 500, err)
            return
        }
        c.r.RenderJSON(w, 200, nil)
    })
}

