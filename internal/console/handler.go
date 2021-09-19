package console

import (
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

type Handler struct {
	logger         logrus.FieldLogger
	userRepository *user.Repository
}

func NewHandler(
	logger logrus.FieldLogger,
	userRepository *user.Repository,
) *Handler {
	return &Handler{
		logger,
		userRepository,
	}
}

func (h *Handler) Migrate(ctx *cli.Context) error {
	userOrm := user.Orm()
	db := h.userRepository.DB()
	db.AutoMigrate(userOrm)
	return nil
}
