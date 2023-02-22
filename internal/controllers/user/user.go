package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/locnguyenvu/mdn/internal/user"
	"github.com/locnguyenvu/mdn/pkg/render"
)

var validate = validator.New()

type Controller struct {
    r *render.Renderer
    userRepository *user.Repository
    logger logrus.FieldLogger
}

func NewController(db *gorm.DB, logger logrus.FieldLogger) *Controller {
    return &Controller{
        render.NewRenderer(logger),
        user.NewRepository(db),
        logger,
    }
}

func loadRequestBody(r *http.Request, target interface{}) error {
    decodeError := json.NewDecoder(r.Body).Decode(target)
    if decodeError != nil {
        return fmt.Errorf("Failed to parse request: %s", decodeError.Error())
    }
    validateErr := validate.Struct(target)
    return validateErr
}

