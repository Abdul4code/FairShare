package validation

import (
	"regexp"
)

// Validator struct holds validation errors
type Validator struct {
	Errors map[string]string
}

// Define regex rule for email
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// New creates and returns a new Validator instance
func New() *Validator {
	return &Validator{
		Errors: map[string]string{},
	}
}

// Valid returns true if there are no validation errors
func (val *Validator) Valid() bool {
	return len(val.Errors) < 1
}

// Add adds a validation error for a given key
func (val *Validator) Add(key string, message string) {
	val.Errors[key] = message
}

// Check adds an error message for the given key if the condition is false
func (val *Validator) Check(ok bool, key string, message string) {
	if !ok {
		val.Add(key, message)
	}
}

// In checks if a value is included in a collection of strings
func (val *Validator) In(value string, collection []string) bool {
	for _, val := range collection {
		if value == val {
			return true
		}
	}
	return false
}

// Is checks if a string matches a given regex pattern
func (val *Validator) Is(value string, pattern *regexp.Regexp) bool {
	if pattern == nil {
		return false
	}
	return pattern.MatchString(value)
}

// Unique checks if all strings in a collection are unique
func (val *Validator) Unique(collections []string) bool {
	uniqueMap := map[string]int{}

	for _, val := range collections {
		uniqueMap[val] = 1
	}

	return len(uniqueMap) <= len(collections)
}
