package model

import (
	"st/backend/db/entity"
	"st/backend/utils/settings"
)

type Metadata map[string]MetadataValue

type MetadataValue struct {
	Id        int64   `json:"id"`
	Value     *string `json:"value"`
	Type      string  `json:"type"`
	UpdatedAt string  `json:"updatedAt"`
	ExpireAt  *string `json:"expireAt"`
}

type UpdateMetadata struct {
	Value    *string `json:"value"`
	Type     *string `json:"type"`
	ExpireAt *string `json:"expireAt"`
}

func (metadata *MetadataValue) ToEntity(key string) *entity.Metadata {
	mr := &entity.Metadata{
		Id:    metadata.Id,
		Key:   key,
		Value: metadata.Value,
		Type:  metadata.Type,
	}

	if metadata.UpdatedAt != "" {
		mr.UpdatedAt = settings.Database.DateFormat
	}

	if metadata.ExpireAt != nil {
		mr.ExpireAt = metadata.ExpireAt
	}

	return mr
}

func (updateMetadata *UpdateMetadata) ToEntity(metadata *entity.Metadata) *entity.Metadata {
	// If Value is an empty string (""), we want to unset it
	// If it is nil we don't want to change it
	// In any other case just change it
	if updateMetadata.Value != nil {
		if *updateMetadata.Value == "" {
			metadata.Value = nil
		} else {
			metadata.Value = updateMetadata.Value
		}
	}

	if updateMetadata.Type != nil {
		metadata.Type = *updateMetadata.Type
	}

	if updateMetadata.ExpireAt == nil {
		metadata.ExpireAt = nil
	} else {
		metadata.ExpireAt = updateMetadata.ExpireAt
	}

	return metadata
}
