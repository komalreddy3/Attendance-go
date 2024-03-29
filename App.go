package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/api"
	"go.uber.org/zap"
)

type App struct {
	router *api.MuxRouter
	logger *zap.SugaredLogger
	db     *pg.DB
}

func NewLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger.Sugar(), err
}

func NewApp(router *api.MuxRouter, logger *zap.SugaredLogger, db *pg.DB) *App {
	app := &App{
		router: router,
		db:     db,
	}
	app.logger = logger
	return app
}

func (app *App) Start() {
	app.router.Init()
	//fmt.Println("running")

	//if err := srv.ListenAndServe(); err != nil {
	//	panic(err)
	//}
	//log.Fatal(http.ListenAndServe("localhost:8080", app.router.Router))
}
