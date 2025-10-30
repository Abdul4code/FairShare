package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Logger struct for logging errors
type Logger struct {
	Log *log.Logger
}

// ErrNotFound is returned when a requested item could not be found
var ErrNotFound = errors.New("the requested Item is not found")

// LogError logs the given data. If panic is true, it panics after logging.
func LogError(data any, panic bool) {
	logger := Logger{
		Log: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile),
	}
	if panic {
		logger.Log.Panic(data)
	}

	logger.Log.Println(data)
}

// WriteJSON writes the given data as a JSON response with the specified status code.
func WriteError(
	w http.ResponseWriter,
	statusCode int,
	data any,
) {
	msg := map[string]any{
		"error": data,
	}

	WriteJSON(w, statusCode, msg)
}

// NotFoundError is a helper function to write a 404 Not Found error response.
func NotFoundError(
	w http.ResponseWriter,
	r *http.Request,
) {
	WriteError(w, http.StatusNotFound, "The requested resource Not found")
}

// MethodNotAllowed is a helper function to write a 405 Method Not Allowed error response.
func MethodNotAllowed(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := fmt.Sprintf("The method %s is not allowed on path %s", r.Method, r.URL)
	WriteError(w, http.StatusMethodNotAllowed, message)
}

// BadRequestError is a helper function to write a 400 Bad Request error response.
func BadRequestError(
	w http.ResponseWriter,
	r *http.Request,
	err any,
) {
	WriteError(w, http.StatusBadRequest, err)
}

// InternalServerError is a helper function to write a 500 Internal Server Error response.
func InternalServerError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	WriteError(w, http.StatusInternalServerError, "Unexpected internal server error")
	LogError(err, false)
}

// UnauthorizedError is a helper function to write a 401 Unauthorized error response.
func UnauthorizedError(
	w http.ResponseWriter,
	r *http.Request,
	err any,
) {
	WriteError(w, http.StatusUnauthorized, "Unauthorized Error: Authentication required")
}

// DuplicateError is a helper function to write a 409 Conflict error response.
func DuplicateError(
	w http.ResponseWriter,
	r *http.Request,
	err any,
) {
	WriteError(w, http.StatusConflict, "This Item already exists")
}
