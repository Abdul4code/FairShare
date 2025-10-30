package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Abdul4code/FairShare/internal"
	"github.com/Abdul4code/FairShare/internal/model"
)

// GroupModel provides database operations for the groups table.
// It holds a reference to a sql.DB connection pool.
type GroupModel struct {
	conn *sql.DB
}

// Insert inserts a new group row into the database and populates
// the given model.Group with the returned id and created_at timestamp.
//
// The function expects the caller to have validated fields on data.
// It returns any error encountered while executing the query or scanning
// the returned row.
func (m GroupModel) Insert(data *model.Group) error {
	query := `INSERT INTO groups (name, currency, description, created_by)
				VALUES ($1, $2, $3, $4)
			  RETURNING id, created_at;
			`

	// QueryRow is used because exactly one row is expected to be returned.
	row := m.conn.QueryRow(
		query,
		data.Name,
		data.Currency,
		data.Description,
		data.CreatedBy,
	)

	// Scan the returned id and created_at into the provided struct.
	// Note: if the schema changes, the RETURNING list must be kept in sync.
	err := row.Scan(&data.Id, &data.CreatedAt)

	return err
}

// Get retrieves a group by its integer id. If the id is invalid (<1)
// ErrNotFound is returned to indicate the resource does not exist.
//
// The function returns a pointer to model.Group on success. It returns
// internal.ErrNotFound when no row is found so callers can distinguish
// "not found" from other errors.
func (m GroupModel) Get(id int) (*model.Group, error) {
	if id < 1 {
		return nil, internal.ErrNotFound
	}

	// Avoid SELECT * in production code because column order matters for Scan.
	// Using explicit column names is safer and clearer.
	query := `SELECT * FROM groups WHERE id = $1;
			`
	row := m.conn.QueryRow(query, id)

	group := &model.Group{}
	err := row.Scan(
		&group.Id,
		&group.Name,
		&group.Currency,
		&group.Description,
		&group.CreatedBy,
		&group.CreatedAt,
		&group.Version,
	)

	// Distinguish between "no rows" and other DB errors so callers can react appropriately.
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, internal.ErrNotFound
		default:
			return nil, err
		}
	}

	return group, err
}

// Update applies changes to an existing group row using optimistic locking
// via the version column. It expects data.Id and data.Version to be set
// to target the correct row/version.
//
// If the id is invalid, ErrNotFound is returned. If the WHERE clause
// matches no rows (concurrent update or missing row), ErrNotFound is returned.
func (m GroupModel) Update(data *model.Group) error {
	// Debug print preserved from original code; remove in production if noisy.
	fmt.Println(data)

	if data.Id < 1 {
		return internal.ErrNotFound
	}

	// Note: RETURNING * will return all columns; the Scan call below MUST match
	// the column order returned by the database. Prefer explicit RETURNING column list
	// to avoid subtle bugs if schema changes.
	query := `UPDATE groups
				SET name = $1,
				currency = $2,
				description = $3,
				version = version + 1
			  WHERE id = $4 AND version=$5
			  RETURNING *
			`
	row := m.conn.QueryRow(
		query,
		data.Name,
		data.Currency,
		data.Description,
		data.Id,
		data.Version,
	)

	// Scan updated row back into the model. Order here must match the DB's column order.
	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.Currency,
		&data.Description,
		&data.CreatedBy,
		&data.CreatedAt,
		&data.Version,
	)

	// If no rows were returned, treat as not found (concurrent edit or deleted).
	if err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return internal.ErrNotFound
		default:
			return err
		}
	}

	return nil
}

// DeleteGroup deletes a group by id. It returns ErrNotFound when the id
// is invalid or when no rows were affected by the delete operation.
//
// Uses Exec instead of QueryRow because no rows are expected to be returned.
func (m *GroupModel) DeleteGroup(id int) error {
	if id < 1 {
		return internal.ErrNotFound
	}

	query := `DELETE FROM groups WHERE id = $1`
	res, err := m.conn.Exec(query, id)
	if err != nil {
		return err
	}

	// RowsAffected tells us whether the delete actually removed a row.
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return internal.ErrNotFound
	}

	return nil
}
