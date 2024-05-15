package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
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
	mux.HandleFunc("DELETE /todo/{id}", h.deleteTodo)

	mux.HandleFunc("GET /todo", h.getTodos)

	mux.HandleFunc("POST /todo", h.addTodo)

	mux.HandleFunc("PUT /todo/{id}", h.updateTodo)
	// END OF API ENDPOINTS
}

func (h *handler) homePage(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}

func (h *handler) addTodo(w http.ResponseWriter, r *http.Request) {
	var id int

	body := &types.Todo{
		CreatedAt:   time.Now(),
		Description: r.FormValue("description"),
		UpdatedAt:   time.Now(),
	}

	query := `
		insert into list
		(createdAt, description, updatedAt)
		values ($1, $2, $3)
		returning id
	`

	err := h.db.QueryRow(
		query,
		body.CreatedAt,
		body.Description,
		body.UpdatedAt,
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error inserting to database")
		return
	}

	views.ToDoCard(body.Description, strconv.Itoa(id)).Render(r.Context(), w)
}

func (h *handler) updateTodo(w http.ResponseWriter, r *http.Request) {
	n, idErr := strconv.Atoi(r.PathValue("id"))

	if idErr != nil {
		fmt.Println(idErr)
		fmt.Println("Invalid id")
		return
	}

	_, err := h.db.Exec(
		"update list set description = $2, updatedAt = $3 where id = $1",
		n,
		r.FormValue("description-update"),
		time.Now(),
	)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Internal server error")
		return
	}

	fmt.Println("Updated")
	h.getTodos(w, r)
}

func (h *handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		fmt.Println(err)
		fmt.Println("Not a valid id")
		return
	}

	result, e := h.db.Query("delete from list where id = $1", n)

	if e != nil {
		fmt.Println(e)
		fmt.Println("Can not delete todo")
		return
	}

	defer func() {
		if err := result.Close(); err != nil {
			fmt.Println(err)
			fmt.Println("Error 500")
			return
		}
	}()

	fmt.Println("Deleting...")

	h.getTodos(w, r)
}

func (h *handler) getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("select * from list order by updatedAt desc")

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
		views.ToDoCard(todo.Description, strconv.Itoa(todo.ID)).Render(r.Context(), w)
	}
}
