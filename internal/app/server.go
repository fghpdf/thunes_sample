package app

import (
	"fghpdf.me/thunes_homework/internal/pkg/ping"
	"fghpdf.me/thunes_homework/internal/pkg/routers"
	"fghpdf.me/thunes_homework/internal/pkg/terminal"
)

func Run() {
	startTerminalApp()
}

func startWebApp()  {
	app := routers.Init()

	router := app.Group("/api")

	router.GET("/connect", ping.Handler)

	app.Run()
}

func startTerminalApp()  {
	app := terminal.Init()

	app.Start()
}
