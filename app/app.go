package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	router *gin.Engine
}

func (a *App) Start(host string) {
	annabingoEnv := os.Getenv("ANNABINGO_ENV")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Error creating logger:", err.Error())
	}
	defer logger.Sync()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(static.Serve("/", static.LocalFile("./ui/build/", false)))
	a.router = router

	db, _ := buntdb.Open("./db/annabingo.db")
	service := NewBingoService(db)
	err = service.CreateIndexOnTitle()
	if err != nil {
		logger.Warn("index on field title not created", zap.Error(err))
	}
	handler := NewBingoHandler(service, logger)
	router.GET("/api", handler.HandleGetTestBingo)
	router.GET("/api/view/:id", handler.HandleGetBingoById)
	router.GET("/api/stats", handler.HandleGetStatistics)
	router.GET("/api/search/:query", handler.HandleSearch)
	router.GET("/api/index", handler.HandleCreateIndex)
	router.POST("/api/create", handler.HandlePostBingo)

	httpServer := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if annabingoEnv == "PROD" {
		cert := "/etc/letsencrypt/live/www.annabingo.de/fullchain.pem"
		key := "/etc/letsencrypt/live/www.annabingo.de/privkey.pem"
		logger.Info("application started over https", zap.String("env", "PROD"))
		err = httpServer.ListenAndServeTLS(cert, key)
	} else {
		logger.Info("application started over http", zap.String("env", "DEV"))
		err = httpServer.ListenAndServe()
	}
	if err != nil {
		logger.Error("application not started",
		 	zap.String("env", annabingoEnv), zap.Error(err))
	}
}
