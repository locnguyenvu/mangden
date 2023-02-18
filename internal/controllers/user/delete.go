package user

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"

)

func (c *Controller) HandleDelete() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64);
        if err != nil {
            c.r.RenderJSON(w, 400, errors.New("Invalid id"))
            return
        }
        if err := c.userRepository.Delete(id); err != nil {
            c.r.RenderJSON(w, 500, err)
            return
        }
        c.r.RenderJSON(w, 200, nil)
    })
}
