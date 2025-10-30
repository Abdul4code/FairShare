package model

import (
	"fmt"

	"github.com/Abdul4code/FairShare/internal/validation"
)

// GroupOutput represents a group object returned to API clients.
type Group struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
	CreatedAt   string `json:"created_at"`
	Version     int    `json:"version"`
}

// GroupInput represents the JSON payload used when creating a group.
type GroupInput struct {
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
}

// GroupUpdate represents the JSON payload used when updating a group.
type GroupUpdate struct {
	Name        *string `json:"name"`
	Currency    *string `json:"currency"`
	Description *string `json:"description"`
	CreatedBy   *int    `json:"created_by"`
}

// Validate checks the GroupInput fields using the provided validation.Validator.
// It returns a map of field -> error message when validation fails, or nil when valid.
func (input *Group) Validate(val *validation.Validator) map[string]string {
	supportedCurrency := []string{"Dollar", "Euro", "Pound", "Naira"}

	val.Check(len(input.Name) > 1, "Name", "The Name of the group cannot be empty")
	val.Check(
		val.In(input.Currency, supportedCurrency),
		"Currency",
		fmt.Sprintf("Unsurported Currency. It should be one of %v", supportedCurrency),
	)

	if ok := val.Valid(); !ok {
		return val.Errors
	}
	return nil
}
