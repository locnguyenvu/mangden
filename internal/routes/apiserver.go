package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"github.com/locnguyenvu/mdn/internal/controllers/user"
)

func APIServer(
    dbconn *gorm.DB,
    logger logrus.FieldLogger,
) http.Handler {

    userController := user.NewController(
        dbconn,
        logger,
    )

    r := chi.NewRouter()
    r.Method("GET", "/users", userController.HandleIndex())
    r.Method("POST", "/users", userController.HandleCreate())
    r.Method("GET", "/users/{id}", userController.HandleGet())
    r.Method("PUT", "/users/{id}", userController.HandleUpdate())
    r.Method("DELETE", "/users/{id}", userController.HandleDelete())
    return r
}
