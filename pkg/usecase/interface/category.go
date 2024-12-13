package interfaces

import (
	"ahava/pkg/domain"
	"ahava/pkg/utils/models"
)

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(categoryID int, category, description string) (domain.Category, error)
	DeleteCategory(categoryID int) error
	GetCategories() ([]domain.Category, error)
	GetProductDetailsInACategory(id int) ([]models.Inventories, error)
	GetBannersForUsers() ([]models.Banner, error)
}
