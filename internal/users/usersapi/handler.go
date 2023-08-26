package usersapi

import (
	"context"
	"github.com/byfood/byfood-core/app"
	"github.com/byfood/byfood-core/internal/uniqid"
	"github.com/byfood/byfood-core/internal/users"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type usersService interface {
	GetAllUsers(ctx context.Context) ([]users.User, error)
	GetUser(ctx context.Context, userID int64) (*users.User, error)
	AddUser(ctx context.Context, u *users.User) error
	UpdateUser(ctx context.Context, u *users.User) error
	DeleteUser(ctx context.Context, u *users.User) error
}

type Handler struct {
	Users usersService
}

type AddUserRequest struct {
	Name string `json:"name"`
}

func (a AddUserRequest) Validate(r *http.Request) error {
	return validation.ValidateStruct(&a, validation.Field(&a.Name, validation.Required))
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddUserRequest
	if !app.BindAndValidate(w, r, &req) {
		return
	}

	var u = users.User{
		ID:   uniqid.Generate(),
		Name: req.Name,
	}

	err := h.Users.AddUser(ctx, &u)
	if err != nil {
		app.InternalError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["id"], 10, 64)

	u, err := h.Users.GetUser(ctx, userID)
	if err != nil {
		app.InternalError(w, r, err)
		return
	}

	app.JSON(w, http.StatusOK, u)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.Users.GetAllUsers(ctx)
	if err != nil {
		app.InternalError(w, r, err)
		return
	}

	app.JSON(w, http.StatusOK, users)
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func (u UpdateUserRequest) Validate(r *http.Request) error {
	return validation.ValidateStruct(&u, validation.Field(&u.Name, validation.Required))
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["id"], 10, 64)

	var req UpdateUserRequest
	if !app.BindAndValidate(w, r, &req) {
		return
	}

	u := users.User{
		ID:   userID,
		Name: req.Name,
	}

	err := h.Users.UpdateUser(ctx, &u)
	if err != nil {
		app.InternalError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["id"], 10, 64)

	var u = users.User{
		ID: userID,
	}

	err := h.Users.DeleteUser(ctx, &u)
	if err != nil {
		app.InternalError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
