package app

import (
	"log"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tidwall/buntdb"
	"go.uber.org/zap"
)

type App struct {
}

func (a *App) Start(host string) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Error creating logger:", err.Error())
	}
	defer logger.Sync()

	router := echo.New()
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Static("/", "ui/build")

	db, _ := buntdb.Open("./db/annabingo.db")
	service := NewBingoService(db)
	err = service.CreateIndexOnTitle()
	if err != nil {
		logger.Warn("index on field title not created", zap.Error(err))
	}
	handler := NewBingoHandler(service, logger)
	router.GET("/api", handler.GetTestBingo)
	router.GET("/api/view/:id", handler.GetBingoById)
	router.GET("/api/stats", handler.GetStatistics)
	router.GET("/api/search/:query", handler.GetSearch)
	router.POST("/api/index", handler.PostCreateIndex)
	router.POST("/api/create", handler.PostBingo)

	if os.Getenv("ENV") == "LIVE" {
		prom := prometheus.NewPrometheus("annabingo", nil)
		prom.Use(router)
	}

	router.Start(":8000")
}
