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
	h.HandleFunc("POST /logout", e.logout)
	h.HandleFunc("POST /register", e.registerUser)
}

func (e *endpoint) logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		MaxAge:   -1,
		Name:     "session_token",
		Path:     "/",
		Value:    "",
	}

	http.SetCookie(w, cookie)
}

func (e *endpoint) registerUser(w http.ResponseWriter, r *http.Request) {
	query := `
		insert into users
		(createdAt, email, name, password, updatedAt)
		values (NOW(), $1, $2, $3, NOW())
		returning id
	`

	pwBytes, pwErr := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)

	if pwErr != nil {
		fmt.Println()
		return
	}

	var id int
	name := r.FormValue("name")

	err := e.db.QueryRow(
		query,
		r.FormValue("email"),
		name,
		string(pwBytes),
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Can not register")
		return
	}

	token, tokenError := utils.CreateJWT(id, name)

	if tokenError != nil {
		fmt.Println(tokenError)
		fmt.Println("Token error")
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		HttpOnly: true,
		Value:    token,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}
