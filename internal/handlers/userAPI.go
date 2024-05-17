package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ernstlegaspi/todolist/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type endpoint struct {
	db *sql.DB
}

func RunUser(db *sql.DB) *endpoint {
	return &endpoint{
		db: db,
	}
}

func (e *endpoint) InitUserEndpoints(h *http.ServeMux) {
	h.HandleFunc("POST /register", e.registerUser)
}

func (e *endpoint) registerUser(w http.ResponseWriter, r *http.Request) {
	query := `
		insert into users
		(createdAt, email, name, password, updatedAt)
		values ($1, $2, $3, $4, $5)
		returning id
	`

	pwBytes, pwErr := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)

	if pwErr != nil {
		fmt.Println()
		return
	}

	var id int

	err := e.db.QueryRow(
		query,
		time.Now(),
		r.FormValue("email"),
		r.FormValue("name"),
		string(pwBytes),
		time.Now(),
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Can not register")
		return
	}

	token, tokenError := utils.CreateJWT(id)

	if tokenError != nil {
		fmt.Println(tokenError)
		fmt.Println("Token error")
		return
	}

	cookie := &http.Cookie{
		Name:  "session_token",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}
