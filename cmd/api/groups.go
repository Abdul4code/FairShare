package main

import (
	"net/http"

	"github.com/Abdul4code/FairShare/internal"
	"github.com/Abdul4code/FairShare/internal/data"
	"github.com/Abdul4code/FairShare/internal/validation"
)

// CreateGroupHandler handles POST /v1/groups. It reads the JSON body into a
// data.GroupInput, validates it and returns either validation errors or the created group.
func (app *application) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	group_info := data.GroupInput{}

	if err := internal.ReadJSON(w, r, &group_info); err != nil {
		internal.WriteJSON(w, 201, err.Error())
		return
	}

	val := validation.New()
	if errors := group_info.Validate(val); errors != nil {
		internal.BadRequestError(w, r, errors)
		return
	}

	internal.WriteJSON(w, http.StatusCreated, group_info)
}
