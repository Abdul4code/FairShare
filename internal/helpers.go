package internal

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Abdul4code/FairShare/internal/validation"
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

// ReadQueryInt reads values from query strings and returns them as integer
func ReadQueryInt(
	r *http.Request,
	val *validation.Validator,
	key string,
	defaultVal int,
) int {
	queries := r.URL.Query()
	value := queries.Get(key)

	// returns rthe default value if key is not found in query
	if value == "" {
		return defaultVal
	}

	// convert the string value to integer
	// add verification error and return default if converstion fails
	int_value, err := strconv.Atoi(value)
	if err != nil {
		val.Add(key, "value must be integer")
		return defaultVal
	}

	return int_value
}

// ReadQueryString reads values from query strings and returns them as string
func ReadQueryString(
	r *http.Request,
	key string,
	defaultVal string,
) string {
	queries := r.URL.Query()
	value := queries.Get(key)

	if value == "" {
		return defaultVal
	}

	return value
}

// ReadQueryString reads values from query strings and returns them as string
func ReadQueryCSV(
	r *http.Request,
	key string,
	defaultVal []string,
) []string {
	queries := r.URL.Query()
	value := queries.Get(key)

	if value == "" {
		return defaultVal
	}

	value_csv := strings.Split(value, ",")

	return value_csv
}

// GetSortValue returns the sort value passed as a query string without its direction
// takes in string ie name- and return name
func GetSortValue(value string) string {
	val := strings.TrimSuffix(value, "-")
	val = strings.TrimSuffix(val, "+")

	return val
}

// GetSortDirection returns the sort direction passed as a query string without its value
// takes in string ie name- and return DESC
func GetSortDirection(value string) string {
	if strings.HasSuffix(value, "-") {
		return "DESC"
	}

	return "ASC"
}
