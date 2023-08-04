package delivery

import (
	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (u *UserHandler) UserAPIRoute(r *mux.Router) error {
	userPath := "/users"
	r.HandleFunc(userPath, u.GetAllUser).Methods(GET)
	r.HandleFunc(userPath+"/{id}", u.GetUserByID).Methods(GET)
	r.HandleFunc(userPath, u.CreateUser).Methods(POST)
	r.HandleFunc(userPath, u.UpdateUser).Methods(PUT)
	r.HandleFunc(userPath, u.UpdateUser).Methods(DELETE)

	return nil
}
