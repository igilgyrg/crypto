package app

import (
	"fmt"
	_ "github.com/igilgyrg/crypto/docs"
	"github.com/igilgyrg/crypto/internal/config"
	"github.com/igilgyrg/crypto/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	swagger "github.com/swaggo/http-swagger"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type App struct {
	cfg    *config.Config
	logger *logging.Logger
	router *httprouter.Router
	server *http.Server
}

func NewApp(cfg *config.Config, logger *logging.Logger) *App {
	logger.Info("route initializing")
	router := httprouter.New()

	logger.Info("swagger router initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", swagger.WrapHandler)
	return &App{
		cfg:    cfg,
		logger: logger,
		router: router,
	}
}

func (a *App) Start() {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if a.cfg.ListenType == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("listen unix socket socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("create unix socket")
		listener, listenErr = net.Listen("unix", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.BindIp, a.cfg.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedOrigins:     []string{"http://localhost:3000"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Authorization", "Location", "Charset", "Access-Control-Allow-Origin", "Content-Type", "content-type", "Origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Access-Token", "Refresh-Token", "Location", "Authorization", "Content-Disposition"},
		Debug:              false,
	})

	handler := c.Handler(a.router)

	server := &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
