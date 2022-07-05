package main

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/config"
	"github.com/alexdyukov/gophermart/internal/gophermart/repository/memory"
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	//config
	appConf := config.NewGophermartConfig()
	// Router
	gophermartRouter := chi.NewRouter()

	// Storage
	//gophermartStore := postgres.NewGophermartStore()

	gophermartStore := memory.NewGophermartStore()

	// Authentication handlers

	// Chi middlewares
	gophermartRouter.Use(middleware.Recoverer)
	// other middlewares, i.e. authorize

	// Handlers
	gophermartRouter.Post("/api/user/orders", handler.PostOrder(usecase.NewLoadOrderNumber(gophermartStore)))
	gophermartRouter.Get("/api/user/orders", handler.GetOrders(usecase.NewListOrderNums(gophermartStore)))
	gophermartRouter.Get("/api/user/balance", handler.GetBalance(usecase.NewShowBalanceState(gophermartStore)))
	gophermartRouter.Post("/api/user/balance/withdraw", handler.PostWithdraw(usecase.NewWithdrawFunds(gophermartStore)))
	gophermartRouter.Get("/api/user/withdrawals", handler.GetWithdrawals(usecase.NewListWithdrawals(gophermartStore))) // BeOl - видать это ошибка

	server := http.Server{
		Addr:    appConf.RunAddr,
		Handler: gophermartRouter,
	}

	err := server.ListenAndServe()
	log.Print(err)
}
