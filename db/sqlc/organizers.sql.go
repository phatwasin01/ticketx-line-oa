// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: organizers.sql

package db

import (
	"context"
	"database/sql"
)

const createOrganizer = `-- name: CreateOrganizer :one
INSERT INTO organizers (
  name,
  email,
  password,
  phone
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, email, password, phone
`

type CreateOrganizerParams struct {
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Phone    sql.NullString `json:"phone"`
}

func (q *Queries) CreateOrganizer(ctx context.Context, arg CreateOrganizerParams) (Organizer, error) {
	row := q.db.QueryRowContext(ctx, createOrganizer,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Phone,
	)
	var i Organizer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Phone,
	)
	return i, err
}

const deleteOrganizer = `-- name: DeleteOrganizer :exec
DELETE FROM organizers WHERE id = $1
`

func (q *Queries) DeleteOrganizer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrganizer, id)
	return err
}

const getOrganizer = `-- name: GetOrganizer :one
SELECT id, name, email, password, phone FROM organizers
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetOrganizer(ctx context.Context, email string) (Organizer, error) {
	row := q.db.QueryRowContext(ctx, getOrganizer, email)
	var i Organizer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Phone,
	)
	return i, err
}

const listOrganizers = `-- name: ListOrganizers :many
SELECT id, name, email, password, phone FROM organizers
ORDER BY id
LIMIT $1 
OFFSET $2
`

type ListOrganizersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrganizers(ctx context.Context, arg ListOrganizersParams) ([]Organizer, error) {
	rows, err := q.db.QueryContext(ctx, listOrganizers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Organizer{}
	for rows.Next() {
		var i Organizer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Phone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrganizerPassword = `-- name: UpdateOrganizerPassword :exec
UPDATE organizers SET password = $2
WHERE id = $1
RETURNING id, name, email, password, phone
`

type UpdateOrganizerPasswordParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdateOrganizerPassword(ctx context.Context, arg UpdateOrganizerPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateOrganizerPassword, arg.ID, arg.Password)
	return err
}
