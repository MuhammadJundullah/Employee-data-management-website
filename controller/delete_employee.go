package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteEmployeeController( db *sql.DB ) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		
			id := r.URL.Query().Get("id")
			r.ParseForm()

			_, err := db.Exec("DELETE FROM employee WHERE id=?", id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte( err.Error()))
				return
			}


			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
	}
}