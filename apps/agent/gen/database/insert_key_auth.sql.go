// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: insert_key_auth.sql

package database

import (
	"context"
)

const insertKeyAuth = `-- name: InsertKeyAuth :exec
INSERT INTO
    ` + "`" + `key_auth` + "`" + ` (
        id,
        workspace_id
    )
VALUES
    (
        ?,
        ?
    )
`

type InsertKeyAuthParams struct {
	ID          string
	WorkspaceID string
}

func (q *Queries) InsertKeyAuth(ctx context.Context, arg InsertKeyAuthParams) error {
	_, err := q.db.ExecContext(ctx, insertKeyAuth, arg.ID, arg.WorkspaceID)
	return err
}
