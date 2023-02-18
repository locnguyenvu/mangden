package user

import (
    "net/http"
    "time"
)

type ListUserResponseObj struct {
    Id int64 `json:"id"`
    Username string `json:"userName"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Yob int `json:"yob"`
    CreatedAt time.Time `json:"createdAt"`
}

func (c *Controller) HandleIndex() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        users, err := c.userRepository.ListLatest()
        if err != nil {
            return
        }
        var userlist []ListUserResponseObj
        for _, elem := range users {
            userlist = append(userlist, ListUserResponseObj{elem.ID, elem.Username, elem.FirstName, elem.LastName, elem.Yob, elem.CreatedAt})
        }
        c.r.RenderJSON(w, 200, userlist)
        return
    })
}
