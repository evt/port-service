package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/evt/port-service/internal/common/errors"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Unauthorised", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Bad request", http.StatusBadRequest)
}

func NotFound(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Not found", http.StatusBadRequest)
}

func RespondWithError(err error, w http.ResponseWriter, r *http.Request) {
	slugError, ok := err.(errors.SlugError)
	if !ok {
		InternalError("internal-server-error", err, w, r)
		return
	}

	switch slugError.ErrorType() {
	case errors.ErrorTypeAuthorization:
		Unauthorised(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeNotFound:
		NotFound(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

func httpRespondWithError(err error, slug string, w http.ResponseWriter, r *http.Request, msg string, status int) {
	log.Printf("error: %s, slug: %s, msg: %s", err, slug, msg)

	resp := ErrorResponse{slug, status}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}
