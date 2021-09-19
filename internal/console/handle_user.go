package console

import (
	"fmt"

	"github.com/locnguyenvu/mangden/internal/user"
	cli "github.com/urfave/cli/v2"
)

func (h *Handler) UserCreate(ctx *cli.Context) error {
	h.logger.Info("Create user command")
	user := &user.User{}
	user.SetUsername("loc")
	user.SetPassword("123456")

	_, err := h.userRepository.Create(user)
	return err
}

func (h *Handler) UserInfo(ctx *cli.Context) error {
	userId := int64(1)
	user, err := h.userRepository.Find(userId)

	fmt.Printf("%s\n", user.Username)
	return err
}

func (h *Handler) UserUpdate(ctx *cli.Context) error {
	userId := int64(1)
	user, err := h.userRepository.Find(userId)
	if err != nil {
		return err
	}
	user.Username = "vulocnguyen"
	return h.userRepository.Update(user)

}
