package internal

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ReadParamId reads the "id" URL parameter from the request and converts it to an integer.
func ReadParamId(r *http.Request) (int, error) {
	params := httprouter.ParamsFromContext(r.Context())
	value := params.ByName("id")

	if value == "" {
		return 0, ErrNotFound
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, ErrNotFound
	}

	return intValue, nil
}
