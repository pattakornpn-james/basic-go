package models

import "time"

type Permissions struct {
	PermissionId int64     `db:"permission_id"`
	DocumentId   int64     `db:"document_id"`
	UserId       int64     `db:"user_id"`
	CanView      bool      `db:"can_view"`
	CanEdit      bool      `db:"can_edit"`
	CanDelete    bool      `db:"can_delete"`
	GrantedAt    time.Time `db:"granted_at"`
}
