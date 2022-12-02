package domain

import (
	"github.com/google/uuid"
	"time"
)

//const VersionPublish = "publish"
//const VersionDraft = "draft"

type ModelUUID struct {
	ID uuid.UUID
}

type ModelBase struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type File struct {
	name string
}

type DraftVersion interface {
	VersionPublish() DraftVersion
}
