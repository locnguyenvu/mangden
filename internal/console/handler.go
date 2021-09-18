package console

import (
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger         logrus.FieldLogger
	userRepository *user.Repository
}

func NewHandler(
	logger logrus.FieldLogger,
	userRepository *user.Repository,
) *Handler {
	logger.Info("Hello world")
	return &Handler{
		logger,
		userRepository,
	}
}
