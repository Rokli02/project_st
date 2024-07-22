package service

import (
	"st/backend/db/repository"
	"st/backend/model"
	"st/backend/utils"
	"st/backend/utils/logger"
)

type MetadataService struct {
	MetaRepo *repository.MetadataRepository
}

func (s *MetadataService) LoadMetadatas() model.Metadata {
	metadataMap := make(model.Metadata)

	metadatas := s.MetaRepo.FindAll()
	for _, metadata := range metadatas {
		metadataMap[metadata.Key] = model.MetadataValue{
			Id:        metadata.Id,
			Value:     metadata.Value,
			Type:      metadata.Type,
			UpdatedAt: utils.ToTime(metadata.UpdatedAt),
		}
	}

	logger.Info("Metadatas are loaded into the memory")

	return metadataMap
}

func (s *MetadataService) UpdateMetadata(id int64, value model.UpdateMetadata) bool {
	// Get metadata from DB by id

	// If doesn't exist return false

	// Otherwise insert into values THEN update

	// If Value or Type is 'nil', then don't change it
	// If empty string (''), then set to 'nil'
	// If any other value, then just simply set it

	return true
}

func (s *MetadataService) CreateMetadata(key string, value model.MetadataValue) *model.MetadataValue {

	return nil
}
