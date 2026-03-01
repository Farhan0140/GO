package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	cnf     *config.Config
	usrRepo repo.UserRepo
}

func NewHandler(
	cnf *config.Config,
	userRepo repo.UserRepo,
) *Handler {
	return &Handler{
		cnf:     cnf,
		usrRepo: userRepo,
	}
}
