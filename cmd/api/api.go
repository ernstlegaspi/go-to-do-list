package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ernstlegaspi/todolist/internal/database"
	"github.com/ernstlegaspi/todolist/internal/handlers"
	"github.com/ernstlegaspi/todolist/internal/utils"
	"github.com/ernstlegaspi/todolist/internal/views"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

	if envError := godotenv.Load("../.env"); envError != nil {
		fmt.Println(envError)
		return envError
	}

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")

		if err == nil {
			if cookie.Value == "" {
				views.Auth().Render(r.Context(), w)
				return
			}

			token, err := utils.ParseJWT(cookie.Value)

			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in parse jwt")
				return
			}

			if !token.Valid {
				fmt.Println("Token is not valid.")
				return
			}

			claims := token.Claims.(jwt.MapClaims)

			views.Home(strings.Split(claims["name"].(string), " ")[0]).Render(r.Context(), w)
			return
		}

		views.Auth().Render(r.Context(), w)
	})

	list := handlers.RunList(db.DB)
	list.InitListEndpoints(router)

	user := handlers.RunUser(db.DB)
	user.InitUserEndpoints(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}
