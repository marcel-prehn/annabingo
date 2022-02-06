package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tidwall/buntdb"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme/autocert"
)

type App struct {
}

func (a *App) Start(host string) {
	//annabingoEnv := os.Getenv("ENV")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Error creating logger:", err.Error())
	}
	defer logger.Sync()

	router := echo.New()
	router.AutoTLSManager.HostPolicy = autocert.HostWhitelist("annabingo.de", "www.annabingo.de")
	router.AutoTLSManager.Email = "marcel.prehn@protonmail.com"
	router.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Static("/", "ui/build")
	//router.Use(static.Serve("/", static.LocalFile("./ui/build/", false)))

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

	router.StartAutoTLS(":443")

	// httpServer := http.Server{
	// 	Addr:    ":8000",
	// 	Handler: router,
	// }

	// if annabingoEnv == "LIVE" {
	// 	cert := "/etc/letsencrypt/live/www.annabingo.de/fullchain.pem"
	// 	key := "/etc/letsencrypt/live/www.annabingo.de/privkey.pem"
	// 	logger.Info("application started over https", zap.String("env", "PROD"))
	// 	err = httpServer.ListenAndServeTLS(cert, key)
	// } else {
	// 	logger.Info("application started over http", zap.String("env", "DEV"))
	// 	err = httpServer.ListenAndServe()
	// }
	// if err != nil {
	// 	logger.Error("application not started",
	// 		zap.String("env", annabingoEnv), zap.Error(err))
	// }
}
