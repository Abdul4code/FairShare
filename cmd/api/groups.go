package main

import (
	"errors"
	"net/http"

	"github.com/Abdul4code/FairShare/internal"
	"github.com/Abdul4code/FairShare/internal/model"
	"github.com/Abdul4code/FairShare/internal/validation"
)

// CreateGroupHandler handles POST /v1/groups. It reads the JSON body into a
// data.GroupInput, validates it and returns either validation errors or the created group.
func (app *application) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupInput := model.GroupInput{}

	if err := internal.ReadJSON(w, r, &groupInput); err != nil {
		internal.BadRequestError(w, r, err.Error())
		return
	}

	group := model.Group{
		Name:        groupInput.Name,
		Currency:    groupInput.Currency,
		Description: groupInput.Description,
		CreatedBy:   groupInput.CreatedBy,
	}

	val := validation.New()
	if errors := group.Validate(val); errors != nil {
		internal.BadRequestError(w, r, errors)
		return
	}

	err := app.Models.Groups.Insert(&group)

	if err != nil {
		internal.InternalServerError(w, r, err)
		return
	}

	internal.WriteJSON(w, http.StatusCreated, group)
}

// GetGroupHandler handles GET /v1/groups/:id. It retrieves the group
// identified by the id URL parameter and returns it as JSON.
func (app *application) GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := internal.ReadParamId(r)

	if err != nil {
		internal.NotFoundError(w, r)
		return
	}

	group, err := app.Models.Groups.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNotFound):
			internal.NotFoundError(w, r)
			return
		default:
			internal.InternalServerError(w, r, err)
			return
		}
	}

	internal.WriteJSON(w, http.StatusOK, group)

}

// GetGroupHandler handles GET /v1/groups/:id. It retrieves the group
// identified by the id URL parameter and returns it as JSON.
func (app *application) UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := internal.ReadParamId(r)
	if err != nil {
		internal.NotFoundError(w, r)
		return
	}

	groupInfo := model.GroupInput{}

	if err := internal.ReadJSON(w, r, &groupInfo); err != nil {
		internal.WriteJSON(w, 201, err.Error())
		return
	}

	group := model.Group{
		Name:        groupInfo.Name,
		Currency:    groupInfo.Currency,
		Description: groupInfo.Description,
		CreatedBy:   groupInfo.CreatedBy,
		Id:          id,
	}

	val := validation.New()
	if errors := group.Validate(val); errors != nil {
		internal.BadRequestError(w, r, errors)
		return
	}

	err = app.Models.Groups.Update(&group)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNotFound):
			internal.NotFoundError(w, r)
			return
		default:
			internal.InternalServerError(w, r, err)
			return
		}
	}

	internal.WriteJSON(w, http.StatusOK, group)
}

// GetGroupHandler handles GET /v1/groups/:id. It retrieves the group
// identified by the id URL parameter and returns it as JSON.
func (app *application) PatchGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := internal.ReadParamId(r)
	if err != nil {
		internal.NotFoundError(w, r)
		return
	}

	groupInput := model.GroupUpdate{}
	err = internal.ReadJSON(w, r, &groupInput)
	if err != nil {
		internal.BadRequestError(w, r, err.Error())
		return
	}

	group, err := app.Models.Groups.Get(id)
	if err != nil {
		internal.NotFoundError(w, r)
		return
	}

	if groupInput.Name != nil {
		group.Name = *groupInput.Name
	}

	if groupInput.Description != nil {
		group.Description = *groupInput.Description
	}

	if groupInput.Currency != nil {
		group.Currency = *groupInput.Currency
	}

	val := validation.New()
	val_errors := group.Validate(val)
	if val_errors != nil {
		internal.BadRequestError(w, r, val_errors)
		return
	}

	err = app.Models.Groups.Update(group)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNotFound):
			internal.NotFoundError(w, r)
			return
		default:
			internal.InternalServerError(w, r, err)
			return
		}
	}

	internal.WriteJSON(w, http.StatusOK, group)
}

// GetGroupHandler handles GET /v1/groups/:id. It retrieves the group
// identified by the id URL parameter and returns it as JSON.
func (app *application) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := internal.ReadParamId(r)
	if err != nil {
		internal.NotFoundError(w, r)
	}

	err = app.Models.Groups.DeleteGroup(id)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNotFound):
			internal.NotFoundError(w, r)
			return
		default:
			internal.InternalServerError(w, r, err)
		}
	}

	message := map[string]string{
		"message": "The item was deleted successfully",
	}
	internal.WriteJSON(w, http.StatusOK, message)
}
