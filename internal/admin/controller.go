package admin

import (
	"github.com/locnguyenvu/mangden/internal/user"
)

type Controller struct {
	userRepository *user.Repository
}

func NewController(
	userRepository *user.Repository,
) *Controller {
	return &Controller{
		userRepository,
	}
}
