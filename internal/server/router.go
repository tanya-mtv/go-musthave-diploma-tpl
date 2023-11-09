package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/handler"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/repository"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/service"
)

func (s *Server) NewRouter(db *sqlx.DB) *gin.Engine {
	authRepo := repository.NewAuthPostgres(db)
	ordersRepo := repository.NewOrdersPostgres(db)
	balanceRepo := repository.NewBalancePostgres(db)

	authService := service.NewAuthStorage(authRepo)
	accountService := service.NewAccountService(balanceRepo)

	h := handler.NewHandler(authService, ordersRepo, accountService, s.cfg, s.log)

	router := gin.New()

	router.POST("/api/user/register", h.SingUp)
	router.POST("/api/user/login", h.SingIn)

	router.POST("/api/user/orders", h.UserIdentify, h.PostOrder)
	router.GET("/api/user/orders", h.UserIdentify, h.GetOrders)

	router.GET("/api/user/balance", h.UserIdentify, h.GetBalance)
	router.POST("/api/user/balance/withdraw", h.UserIdentify, h.Withdraw)
	router.GET("/api/user/withdrawals", h.UserIdentify, h.GetWithdraws)

	return router
}
