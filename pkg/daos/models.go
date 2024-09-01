// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package daos

import (
	"database/sql"
)

type Artifact struct {
	ID            string
	Content       []byte
	ContentSha256 string
	Type          string
	LastModified  int64
	CollectionID  sql.NullString
}

type Collection struct {
	ID             string
	Name           string
	IsFolderScoped int64
	FolderHash     sql.NullString
	LastModified   int64
}