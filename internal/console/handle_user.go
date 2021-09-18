package console

import (
	cli "github.com/urfave/cli/v2"
)

func (h *Handler) UserCreate(ctx *cli.Context) error {
	h.logger.Info("Create user command")
	return nil
}
