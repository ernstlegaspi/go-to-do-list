package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ernstlegaspi/todolist/internal/types"
	"github.com/ernstlegaspi/todolist/internal/views"
)

type handler struct {
	db *sql.DB
}

func Run(db *sql.DB) *handler {
	return &handler{
		db: db,
	}
}

func (h *handler) InitListEndpoints(mux *http.ServeMux) {
	// START OF PAGES ENDPOINT
	mux.HandleFunc("/", h.homePage)
	// END OF PAGES ENDPOINT

	// START OF API ENDPOINTS
	mux.HandleFunc("GET /todo", h.getTodos)

	mux.HandleFunc("POST /todo", h.addTodo)
	// END OF API ENDPOINTS
}

func (h *handler) homePage(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
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

	_, err := h.db.Exec(
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

	views.ToDoCard(body.Description).Render(r.Context(), w)
}

func (h *handler) getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("select * from list")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Can not fetch todos")
		return
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
			return
		}

		todos = append(todos, todo)
	}

	for _, todo := range todos {
		views.ToDoCard(todo.Description).Render(r.Context(), w)
	}
}
