package console

import (
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

type Handler struct {
	logger              logrus.FieldLogger
	userRepository      *user.Repository
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
	db := h.userRepository.DB()
        err := db.Migrator().AutoMigrate(
            config.AppConfigOrm(),
            user.Orm(), 
        )
        if err != nil {
            h.logger.Error(err)
        }
	return nil
}
