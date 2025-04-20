package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewUpdateEmployeeController( db *sql.DB ) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]
			_, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES(?, ?, ?)", name, npwp, address)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)

			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT name, npwp, address FROM employee WHERE id = ?", id)
			if row.Err() != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			var employee Employee
			err := row.Scan(
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)

			employee.Id = id

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			
			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}
