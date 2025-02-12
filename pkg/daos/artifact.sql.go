// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: artifact.sql

package daos

import (
	"context"
	"database/sql"
)

const deleteArtifactById = `-- name: DeleteArtifactById :exec
DELETE FROM artifacts WHERE id = ?1
`

func (q *Queries) DeleteArtifactById(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteArtifactById, id)
	return err
}

const getMostRecentArtifact = `-- name: GetMostRecentArtifact :one
SELECT id, content, content_sha256, last_modified, collection_id FROM artifacts ORDER BY last_modified DESC LIMIT 1
`

func (q *Queries) GetMostRecentArtifact(ctx context.Context) (Artifact, error) {
	row := q.db.QueryRowContext(ctx, getMostRecentArtifact)
	var i Artifact
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.ContentSha256,
		&i.LastModified,
		&i.CollectionID,
	)
	return i, err
}

const updateArtifactLastModified = `-- name: UpdateArtifactLastModified :exec
UPDATE artifacts SET last_modified = unixepoch() WHERE id = ?1
`

func (q *Queries) UpdateArtifactLastModified(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updateArtifactLastModified, id)
	return err
}

const upsertArtifact = `-- name: UpsertArtifact :one
INSERT INTO artifacts (content, content_sha256, collection_id) VALUES (?1, sha256(?1), ?2)
  ON CONFLICT(content_sha256, COALESCE(collection_id, 'root'))
  DO UPDATE SET last_modified = unixepoch()
  RETURNING id, content, content_sha256, last_modified, collection_id
`

type UpsertArtifactParams struct {
	Content      []byte
	CollectionID sql.NullString
}

func (q *Queries) UpsertArtifact(ctx context.Context, arg UpsertArtifactParams) (Artifact, error) {
	row := q.db.QueryRowContext(ctx, upsertArtifact, arg.Content, arg.CollectionID)
	var i Artifact
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.ContentSha256,
		&i.LastModified,
		&i.CollectionID,
	)
	return i, err
}
