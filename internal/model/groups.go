package model

import (
	"fmt"
	"strings"

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

// GroupQuery represents the JSON item created from request query parameters
type GroupQuery struct {
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	Sort        string `json:"sort"`
}

// ValidateGroupQuery checks the GroupQuery fields using the provided validation.Validator.
// It returns a map of field -> error message when validation fails, or nil when valid.
func (input *GroupQuery) ValidateGroupQuery(val *validation.Validator) map[string]string {
	supportedCurrency := []string{"Dollar", "Euro", "Pound", "Naira"}
	supportedSortFields := []string{"name", "currency", "created_at", "id"}

	// check that the currency is among the list of supported currencies
	val.Check(
		val.In(input.Currency, supportedCurrency) || input.Currency == "",
		"currency",
		fmt.Sprintf("Unsurported Currency. It should be one of %v", supportedCurrency),
	)

	// check that page is a value between 1 to 10,000,000
	val.Check(
		input.Page >= 1 && input.Page <= 10_000_000,
		"page",
		"unsurported Page Value. Value should be between 1 and 10,000,000",
	)

	// check that limit is a value between 1 and 100
	val.Check(
		input.PageSize >= 1 && input.PageSize <= 100,
		"page_size",
		"unsurported page size value: Value should be between 1 and 100",
	)

	// check that the sort is in the possible values to sort by
	sort := strings.TrimSuffix(input.Sort, "-")
	sort = strings.TrimSuffix(sort, "+")

	val.Check(
		val.In(sort, supportedSortFields) || input.Sort == "",
		"sort",
		fmt.Sprintf("Unsurported Sort values. It should be one of %v", supportedSortFields),
	)

	if ok := val.Valid(); !ok {
		return val.Errors
	}
	return nil
}

// Validate checks the Group fields using the provided validation.Validator.
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
