package main

import (
	config "SecondAssignment/config"
	"SecondAssignment/database"
	"SecondAssignment/service/domain/user/delivery"
	"SecondAssignment/service/domain/user/usecase"
	repo2 "SecondAssignment/service/repo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	schema := config.NewSchema()
	db, err := database.Open(*schema)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo2.NewUserRepo(db)
	userUseCase := usecase.NewUserUseCase(repo)
	userHandler := delivery.NewUserHandler(userUseCase)

	r := mux.NewRouter()
	_ = userHandler.UserAPIRoute(r)
	addr := ":3000"
	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
