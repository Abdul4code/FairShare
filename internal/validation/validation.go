package validation

import (
	"regexp"
)

type Validator struct {
	Errors map[string]string
}

// Define regex rule for email
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func New() *Validator {
	return &Validator{
		Errors: map[string]string{},
	}
}

func (val *Validator) Valid() bool {
	return len(val.Errors) < 1
}

func (val *Validator) Add(key string, message string) {
	val.Errors[key] = message
}

func (val *Validator) Check(ok bool, key string, message string) {
	if !ok {
		val.Add(key, message)
	}
}

func (val *Validator) In(value string, collection []string) bool {
	for _, val := range collection {
		if value == val {
			return true
		}
	}
	return false
}

func (val *Validator) Is(value string, pattern *regexp.Regexp) bool {
	if pattern == nil {
		return false
	}
	return pattern.MatchString(value)
}

func (val *Validator) Unique(collections []string) bool {
	uniqueMap := map[string]int{}

	for _, val := range collections {
		uniqueMap[val] = 1
	}

	return len(uniqueMap) <= len(collections)
}
