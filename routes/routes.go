package routes

import (
	"database/sql"
	"go-crud-employee/controller"
	"net/http"
)

func MapRoutes( server *http.ServeMux, db *sql.DB ) {
	server.HandleFunc("/", controller.NewHelloworldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}