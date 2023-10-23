package handler

import (
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/config"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/logger"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/repository"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/service"
)

type Handler struct {
	authService    service.Autorisation
	ordersService  repository.Orders
	accountService service.Account
	cfg            *config.ConfigServer
	log            logger.Logger
}

func NewHandler(auth service.Autorisation, orders repository.Orders, account service.Account, cfg *config.ConfigServer, log logger.Logger) *Handler {
	return &Handler{
		authService:    auth,
		ordersService:  orders,
		accountService: account,
		cfg:            cfg,
		log:            log,
	}
}
