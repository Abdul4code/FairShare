package internal

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

// Logger type holds the standard logger item for this project
type Logger struct {
	Log zerolog.Logger
}

// ErrNotFound is returned when a requested item could not be found
var ErrNotFound = errors.New("the requested Item is not found")

// LogError writes the given data to stdout (using zerolog) and appends a
// JSON Lines (jsonl) entry to errors.jsonl in the repository root.
// The jsonl entry includes a UTC timestamp, the formatted error string and
// the panic flag so automated tooling can consume the file.
//
// Note: we keep writing to stdout with zerolog as a fallback if file IO fails.
func NewLogger() *Logger {
	L := &Logger{}
	L.Log = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	f, ferr := os.OpenFile("errors.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if ferr != nil {
		// if opening the file fails, record that fact to stdout logger
		L.Log.Error().Err(ferr).Msg("failed to open errors.jsonl for appending")
		return L
	}

	multi := zerolog.MultiLevelWriter(L.Log, f)
	multi_logger := zerolog.New(multi).With().Timestamp().Caller().Logger()

	return &Logger{
		Log: multi_logger,
	}
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
	logger := NewLogger()
	logger.Log.Error().Err(err).Msg("internal server error")
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
