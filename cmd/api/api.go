package api

import (
	"fmt"
	"net/http"

	"github.com/ernstlegaspi/todolist/internal/database"
	"github.com/ernstlegaspi/todolist/internal/handlers"
)

type server struct {
	addr string
}

func InitServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) RunAPI() error {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("../internal/static"))

	db, err := database.ConnectDB()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating db")
		return err
	}

	if listTableErr := db.CreateTables(); listTableErr != nil {
		fmt.Println(listTableErr)
		fmt.Println("Error create tables")
		return listTableErr
	}

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	list := handlers.Run(db.DB)
	list.InitListEndpoints(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}
