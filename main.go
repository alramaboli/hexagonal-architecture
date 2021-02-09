package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/projects/hexagonal-architecture/api"
	"github.com/projects/hexagonal-architecture/config"
	"github.com/projects/hexagonal-architecture/domain"
	"github.com/projects/hexagonal-architecture/repository"
)

func main() {

	/*
		mongoURL := "mongodb://localhost:27017"
		mongodb := "product"
		timeout := 5
	*/
	conf, _ := config.NewConfig("./config/config.yaml")
	repo, _ := repository.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	service := domain.NewProductService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products/{code}", handler.Get)
	r.Post("/products", handler.Post)
	r.Delete("/products/{code}", handler.Delete)
	r.Get("/products", handler.GetAll)
	log.Fatal(http.ListenAndServe(conf.Server.Port, r))

}
