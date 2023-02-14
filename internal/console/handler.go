package console

import (
	appconfig "github.com/locnguyenvu/mangden/internal/config"
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

type Handler struct {
	logger              logrus.FieldLogger
	appconfigRepository *appconfig.Repository
	userRepository      *user.Repository
}

func NewHandler(
	logger logrus.FieldLogger,
	appconfigRepository *appconfig.Repository,
	userRepository *user.Repository,
) *Handler {
	return &Handler{
		logger,
		appconfigRepository,
		userRepository,
	}
}

func (h *Handler) Migrate(ctx *cli.Context) error {
	db := h.userRepository.DB()
        err := db.Migrator().AutoMigrate(user.Orm(), appconfig.Orm())
        if err != nil {
            h.logger.Error(err)
        }
	return nil
}
