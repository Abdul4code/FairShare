package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// WriteJSON writes the provided data as JSON to the http.ResponseWriter with the given
// statusCode. It marshals the value, sets the Content-Type header and writes the response.
// On marshal or write failure the error is logged via LogError.
func WriteJSON(
	w http.ResponseWriter,
	statusCode int,
	data any,
) {
	obj, err := json.Marshal(data)
	if err != nil {
		logger := NewLogger()
		logger.Log.Panic().Err(err).Msg("failed to marshal JSON")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err = w.Write(obj); err != nil {
		logger := NewLogger()
		logger.Log.Panic().Err(err).Msg("failed to write JSON response")
		return
	}
}

// ReadJSON decodes JSON from the provided http.Request into data. It enforces a 1MB
// body limit, disallows unknown fields and returns detailed errors for malformed or
// invalid request bodies.
func ReadJSON(
	w http.ResponseWriter,
	r *http.Request,
	data any,
) error {
	// limit the maximum request body size to 1mb
	maxSizeLimit := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxSizeLimit))

	// setup a new json decoder
	dec := json.NewDecoder(r.Body)

	// set the decoder to raise error when there is an unknown field
	dec.DisallowUnknownFields()

	// read the data from the decoder
	err := dec.Decode(data)

	if err != nil {
		var SyntaxError *json.SyntaxError
		var UnmarshalTypeError *json.UnmarshalTypeError
		var InvalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &SyntaxError):
			return fmt.Errorf("invalid request body. Syntax error on line %d", SyntaxError.Offset)
		case errors.As(err, &UnmarshalTypeError):
			if UnmarshalTypeError.Field != "" {
				return fmt.Errorf("invalid Request body. The JSON object has invalid type %v", UnmarshalTypeError.Field)
			} else {
				return errors.New("invalid request body. Syntax Error")
			}
		case errors.Is(err, io.EOF):
			return fmt.Errorf("invalid Request body: The request body cannot be empty")
		case errors.As(err, &InvalidUnmarshalError):
			logger := NewLogger()
			logger.Log.Panic().Err(err).Msg("internal Server error: Failed to read JSON body")

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("invalid Request body: Sysntax error")

		case strings.Contains(err.Error(), "too large"):
			return errors.New("invalid request body: The request body cannot be larger than 1mb")
		case strings.Contains(err.Error(), "unknown field"):
			return fmt.Errorf("invalid request body: the request contains unknown field %v", strings.TrimPrefix(err.Error(), "json: unknown field "))
		}
	}

	err = dec.Decode(&struct{}{})

	if !errors.Is(err, io.EOF) {
		return errors.New("invalid request body: request contain multiple JSON")
	}

	return nil
}
