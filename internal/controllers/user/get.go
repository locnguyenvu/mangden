
package user

import (
    "errors"
    "net/http"
    "strconv"
    "time"

	"github.com/go-chi/chi/v5"
)

type GetUserResponseObj struct {
    Id int64 `json:"id"`
    Username string `json:"userName"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Yob int `json:"yob"`
    CreatedAt time.Time `json:"createdAt"`
}

func (c *Controller) HandleGet() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64);
        if err != nil {
            c.r.RenderJSON(w, 400, errors.New("Invalid id"))
            return
        }
        u := c.userRepository.Get(id)
        if u == nil {
            c.r.RenderJSON(w, 404, errors.New("User not found"))
            return
        }
        c.r.RenderJSON(w, 200, map[string]interface{}{
            "data": GetUserResponseObj{u.ID, u.Username, u.FirstName, u.LastName, u.Yob, u.CreatedAt},
        })
        return
    })
}
