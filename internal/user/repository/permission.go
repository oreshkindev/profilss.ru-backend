package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type PermissionRepository struct {
	database *database.Database
}

func NewPermissionRepository(database *database.Database) *PermissionRepository {
	return &PermissionRepository{database: database}
}

func (repository *PermissionRepository) Get(title string) (*entity.Permission, error) {
	entity := entity.Permission{}

	return &entity, repository.database.Where("id = ?", title).First(&entity).Error
}
