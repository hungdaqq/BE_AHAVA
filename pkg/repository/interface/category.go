package interfaces

import (
	"ahava/pkg/domain"
	"ahava/pkg/utils/models"
)

type CategoryRepository interface {
	AddCategory(category domain.Category) (domain.Category, error)
	CheckCategory(currrent string) (bool, error)
	UpdateCategory(categoryID int, category, description string) (domain.Category, error)
	DeleteCategory(categoryID int) error
	GetCategories() ([]domain.Category, error)
	GetBannersForUsers() ([]models.Banner, error)
	GetImagesOfProductsFromACategory(CategoryID int) ([]string, error)
}
