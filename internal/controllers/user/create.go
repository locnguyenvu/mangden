package user

import "net/http"

type CreateUserRequestBody struct {
    Username string `json:"userName" validate:"required"`
    Firstname string `json:"firstName" validate:"required"`
    Lastname string `json:"lastName" validate:"required"`
    Password string `json:"password" validate:"required"`
    Yob int `json:"yob" validate:"required"`
}

func (c *Controller) HandleCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        payload := CreateUserRequestBody{}
        if err := loadRequestBody(r, &payload); err != nil {
            c.r.RenderJSON(w, http.StatusBadRequest, err)
            return
        }
        _, err := c.userRepository.Create(
            payload.Username,
            payload.Password,
            payload.Firstname,
            payload.Lastname,
            payload.Yob)
        if err != nil {
            c.r.RenderJSON(w, http.StatusBadRequest, err)
            return
        }
        c.r.RenderJSON(w, http.StatusOK, nil)
    })
}
