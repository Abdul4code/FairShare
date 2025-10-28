package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Logger struct {
	Log *log.Logger
}

func LogError(data any, panic bool) {
	logger := Logger{
		Log: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile),
	}
	if panic {
		logger.Log.Panic(data)
	}

	logger.Log.Println(data)
}

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

func NotFoundError(
	w http.ResponseWriter,
	r *http.Request,
) {
	WriteError(w, http.StatusNotFound, "The requested resource Not found")
}

func MethodNotAllowed(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := fmt.Sprintf("The method %s is not allowed on path %s", r.Method, r.URL)
	WriteError(w, http.StatusMethodNotAllowed, message)
}

func BadRequestError(
	w http.ResponseWriter,
	r *http.Request,
	err any,
) {
	WriteError(w, http.StatusBadRequest, err)
}

func InternalServerError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	WriteError(w, http.StatusInternalServerError, "Unexpected internal server error")
	LogError(err, false)
}

func UnauthorizedError(
	w http.ResponseWriter,
	r *http.Request,
	err any,
) {
	WriteError(w, http.StatusInternalServerError, "Unauthorized Error: Authentication required")
}

func DuplicateError(
	w http.ResponseWriter,
	r *http.Request,
	err any,
) {
	WriteError(w, http.StatusConflict, "This Item already exists")
}
