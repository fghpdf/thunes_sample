package app

import (
	"fghpdf.me/thunes_homework/internal/pkg/config"
	"fghpdf.me/thunes_homework/internal/pkg/country"
	"fghpdf.me/thunes_homework/internal/pkg/payer"
	"fghpdf.me/thunes_homework/internal/pkg/ping"
	"fghpdf.me/thunes_homework/internal/pkg/quotation"
	"fghpdf.me/thunes_homework/internal/pkg/routers"
	"fghpdf.me/thunes_homework/internal/pkg/transaction"
	"log"
)

func Run() {
	startWebApp()
}

func startWebApp() {
	config.Init()
	app := routers.Init()

	router := app.Group("/api")

	router.GET("/connect", ping.Handler)
	router.GET("/countries", country.List)
	router.GET("/payers", payer.List)
	router.POST("/quotation", quotation.Create)
	router.POST("/quotations/:id/transactions", transaction.Create)
	router.POST("/transactions/:id/confirm", transaction.Confirm)

	err := app.Run()
	if err != nil {
		log.Fatalf("App start error: %v\n", err)
	}
}
