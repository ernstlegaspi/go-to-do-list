package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	// "time"

	// "golang.org/x/crypto/bcrypt"
	"github.com/ernstlegaspi/todolist/internal/utils"
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
	token, _ := utils.CreateJWT(22)

	fmt.Println(token)
	// query := `
	// 	insert into users
	// 	(createdAt, email, name, password, updatedAt)
	// 	values ($1, $2, $3, $4, $5)
	// `

	// pwBytes, pwErr := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)

	// if pwErr != nil {
	// 	fmt.Println()
	// 	return
	// }

	// _, err := e.db.Exec(
	// 	query,
	// 	time.Now(),
	// 	r.FormValue("email"),
	// 	r.FormValue("name"),
	// 	string(pwBytes),
	// 	time.Now(),
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("Can not register")
	// 	return
	// }

}
