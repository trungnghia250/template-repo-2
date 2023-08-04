package delivery

import (
	"SecondAssignment/service/domain/user/usecase"
	"SecondAssignment/service/model/dto"
	"encoding/json"
	"github.com/gorilla/mux"

	"net/http"
)

type UserHandler struct {
	user usecase.UserUseCase
}

func NewUserHandler(user usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		user: user,
	}
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}

func (u *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.user.GetAllUser(r.Context())
	if err != nil {
		JSON(w, http.StatusBadRequest, err)
	}

	JSON(w, http.StatusOK, users)
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := u.user.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, user)
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	er1 := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	er2 := u.user.CreateUser(r.Context(), req)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, nil)
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserRequest
	er1 := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	er2 := u.user.UpdateUser(r.Context(), req)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, nil)
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	err := u.user.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, nil)
}
