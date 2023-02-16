package testserver

import (
    "net/http"
    "encoding/json"

	"gorm.io/gorm"
    "mck.co/fuel/internal/user"
)

type Handler struct {
    db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
    return &Handler{db}
}

func (h Handler) Migrate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    err := h.db.AutoMigrate(&user.User{})
    if err != nil {
        data := map[string]interface{}{
            "status": "error",
            "message": err.Error(),
        }
        jData, _ := json.Marshal(data)
        w.Write(jData)
        return
    }
    data := map[string]interface{}{
        "status": "ok",
    }
    jData, _ := json.Marshal(data)
    w.Write(jData)
}


type CreateUserRequest struct {
    Username string
    Password string
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    requestBody := CreateUserRequest{}
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        data := map[string]interface{}{
            "status": "error",
            "message": err.Error(),
        }
        jData, _ := json.Marshal(data)
        w.Write(jData)
        return
    }

    user := &user.User{
        Username: requestBody.Username,
    }
    user.SetPassword(requestBody.Password)
    h.db.Save(user)

    data := map[string]interface{}{
        "status": "ok",
    }
    jData, _ := json.Marshal(data)
    w.Write(jData)
}

func (h Handler) ListUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    users := new([]user.User)
    h.db.Find(&users)

    data := map[string]interface{}{
        "status": "ok",
        "data": users,
    }
    jData, _ := json.Marshal(data)
    w.Write(jData)
}

