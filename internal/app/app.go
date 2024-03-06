package app

import (
	"fmt"
	"net/http"

	"github.com/KadyrPoyraz/httplayout/config"
	projectshandler "github.com/KadyrPoyraz/httplayout/internal/handler/http/projects"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
	"github.com/KadyrPoyraz/httplayout/internal/repository/projects"
	projectsService "github.com/KadyrPoyraz/httplayout/internal/service/projects"
	"github.com/gorilla/mux"
)

type app struct{
    cnf config.Config
}

func New(cnf config.Config) App {
	return &app{
        cnf: cnf,
    }
}

func (a *app) Run() error {
    mainRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

    dsnFormat := "postgresql://%s:%s@%s:%s/%s?sslmode=disable"
    dbCnf := a.cnf.DB
    dsn := fmt.Sprintf(dsnFormat, dbCnf.User, dbCnf.Password, dbCnf.Host, dbCnf.Port, dbCnf.Name)
    db, err := db.NewPostgresqlDB(dsn)
    if err != nil {
        return err
    }

    projectsRepo := projects.NewPostgresqlRepo(db)
    projectsService := projectsService.New(db, projectsRepo)
    projectshandler.New(mainRouter, projectsService)

    port := ":" + a.cnf.App.Port
    fmt.Printf("Server started on localhost%s\n", port)
    err = http.ListenAndServe(port, mainRouter)
    if err != nil {
        panic(err)
    }

    return nil
}

