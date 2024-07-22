package model

import (
	"st/backend/db/entity"
	"st/backend/utils"
	"st/backend/utils/settings"
	"time"
)

type Metadata map[string]MetadataValue

type MetadataValue struct {
	Id        int64
	Value     *string
	Type      string
	UpdatedAt *time.Time
	ExpireAt  *time.Time
}

func (m *MetadataValue) ToEntity(key string) *entity.Metadata {
	mr := &entity.Metadata{
		Id:    m.Id,
		Key:   key,
		Value: m.Value,
		Type:  m.Type,
	}

	if m.UpdatedAt != nil {
		mr.UpdatedAt = m.UpdatedAt.Format(settings.Database.DateFormat)
	}

	if m.ExpireAt != nil {
		mr.ExpireAt = utils.ToRef(m.ExpireAt.Format(settings.Database.DateFormat))
	}

	return mr
}

type UpdateMetadata struct {
	Value    *string
	Type     *string
	ExpireAt *time.Time
}

func (m *UpdateMetadata) ToEntity(metadata *entity.Metadata) *entity.Metadata {
	// If Value is an empty string (""), we want to unset it
	// If it is nil we don't want to change it
	// In any other case just change it
	if m.Value != nil {
		if *m.Value == "" {
			metadata.Value = nil
		} else {
			metadata.Value = m.Value
		}
	}

	if m.Type != nil {
		metadata.Type = *m.Type
	}

	if m.ExpireAt == nil {
		metadata.ExpireAt = nil
	} else {
		metadata.ExpireAt = utils.ToRef(m.ExpireAt.Format(settings.Database.DateFormat))
	}

	return metadata
}
