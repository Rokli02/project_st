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

func (s *MetadataService) UpdateMetadata(id int64, updateMetadata *model.UpdateMetadata) *model.MetadataValue {
	// Get metadata from DB by id, if doesn't exist return false
	metadata := s.MetaRepo.FindById(id)
	if metadata == nil {
		return nil
	}

	// Otherwise insert updateMetadata values into metadata, THEN update
	metadata = updateMetadata.ToEntity(metadata)

	if s.MetaRepo.UpdateOne(id, metadata) {
		metadataValue := &model.MetadataValue{
			Id:        metadata.Id,
			Value:     metadata.Value,
			Type:      metadata.Type,
			UpdatedAt: utils.ToTime(metadata.UpdatedAt),
		}

		if metadata.ExpireAt != nil {
			metadataValue.ExpireAt = utils.ToTime(*metadata.ExpireAt)
		}

		return metadataValue
	}

	return nil
}

func (s *MetadataService) CreateMetadata(key string, metadata *model.MetadataValue) *model.MetadataValue {
	if s.MetaRepo.IsExist(key) {
		logger.WarningF("Metadata with such key (%s) is already in use", key)

		return nil
	}

	id := s.MetaRepo.InsertOne(metadata.ToEntity(key))

	if id < 1 {
		return nil
	}

	metadataFromDB := s.MetaRepo.FindById(id)
	if metadataFromDB == nil {
		logger.Warning("Metadata wasn't inserted into database")

		return nil
	}

	newMetadata := &model.MetadataValue{
		Id:        metadataFromDB.Id,
		Value:     metadataFromDB.Value,
		Type:      metadataFromDB.Type,
		UpdatedAt: utils.ToTime(metadataFromDB.UpdatedAt),
	}

	if metadataFromDB.ExpireAt != nil {
		newMetadata.ExpireAt = utils.ToTime(*metadataFromDB.ExpireAt)
	}

	return newMetadata
}
