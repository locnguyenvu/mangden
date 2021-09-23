package console

import cli "github.com/urfave/cli/v2"

func (h *Handler) ConfigSet(ctx *cli.Context) error {
	name := ctx.String("name")
	value := ctx.String("value")

	return h.appconfigRepository.NewOrUpdate(name, value)
}

func (h *Handler) ConfigGet(ctx *cli.Context) error {
	name := ctx.String("name")

	h.logger.Info(h.appconfigRepository.Get(name))
	return nil
}
