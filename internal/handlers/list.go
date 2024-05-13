package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/ernstlegaspi/todolist/internal/types"
)

type handler struct {
	db *sql.DB
}

func Run(db *sql.DB) *handler {
	return &handler{
		db: db,
	}
}

func (h *handler) InitEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/", h.homePage)
}

func (h *handler) homePage(w http.ResponseWriter, r *http.Request) {
	todos, err := h.getTodos()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Can not get todos")
		return
	}

	if r.URL.Path == "/home.js" {
		w.Header().Set("Content-Type", "application/javascript")
		http.ServeFile(w, r, "../internal/views/home/home.js")
		return
	}

	templ := template.Must(template.ParseFiles("../internal/views/home/Home.html"))
	templ.Execute(w, map[string][]*types.Todo{
		"Todos": todos,
	})
}

func (h *handler) addTodo(w http.ResponseWriter, r *http.Request) {
	body := &types.Todo{
		CreatedAt:   time.Now(),
		Description: r.FormValue("description"),
		UpdatedAt:   time.Now(),
	}

	query := `
		insert into list
		(createdAt, description, updatedAt)
		values ($1, $2, $3)
	`

	_, err := h.db.Query(
		query,
		body.CreatedAt,
		body.Description,
		body.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error inserting to database")
		return
	}

	fmt.Println("List inserted.")
}

func (h *handler) getTodos() ([]*types.Todo, error) {
	rows, err := h.db.Query("select * from list")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Can not fetch todos")
		return nil, err
	}

	defer rows.Close()

	todos := []*types.Todo{}

	for rows.Next() {
		todo := new(types.Todo)

		err := rows.Scan(
			&todo.ID,
			&todo.CreatedAt,
			&todo.Description,
			&todo.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Error fetching todo")
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
