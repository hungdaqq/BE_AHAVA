package repository

import (
	"ahava/pkg/domain"
	interfaces "ahava/pkg/repository/interface"
	"ahava/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{DB}
}

func (p *categoryRepository) AddCategory(c domain.Category) (domain.Category, error) {
	var b int
	err := p.DB.Raw(`INSERT INTO categories (category, description) VALUES (?, ?) RETURNING id`,
		c.Category, c.Description).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}

	var categoryResponse domain.Category
	err = p.DB.Raw(`SELECT * FROM categories p WHERE p.id = ?`, b).Scan(&categoryResponse).Error

	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil

}

func (p *categoryRepository) CheckCategory(current string) (bool, error) {
	var i int
	err := p.DB.Raw("SELECT COUNT(*) FROM categories WHERE category=?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}

	if i == 0 {
		return false, err
	}

	return true, err
}

func (p *categoryRepository) UpdateCategory(categoryID int, category, description string) (domain.Category, error) {
	var updateCategory domain.Category

	// Update the category
	if err := p.DB.Raw("UPDATE categories SET category = $1, description = $2 WHERE id = $3 RETURNING *",
		category, description, categoryID).Scan(&updateCategory).Error; err != nil {
		return domain.Category{}, err
	}

	return updateCategory, nil
}

func (c *categoryRepository) DeleteCategory(categoryID int) error {

	result := c.DB.Exec("DELETE FROM categories WHERE id = ?", categoryID)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (c *categoryRepository) GetCategories() ([]domain.Category, error) {
	var model []domain.Category
	err := c.DB.Raw("SELECT * FROM categories").Scan(&model).Error
	if err != nil {
		return []domain.Category{}, err
	}

	return model, nil
}

func (c *categoryRepository) GetBannersForUsers() ([]models.Banner, error) {
	var banners []models.Banner
	err := c.DB.Raw(`select offers.category_id,categories.category as category_name,offers.discount_rate as discount_percentage
	 from offers
	 join categories on categories.id = offers.category_id
	 where offers.discount_rate > 10 
	 Order by offers.discount_rate desc
	 limit 3`).Scan(&banners).Error
	if err != nil {
		return []models.Banner{}, err
	}
	return banners, nil
}

func (c *categoryRepository) GetImagesOfProductsFromACategory(CategoryID int) ([]string, error) {
	var images []string
	err := c.DB.Raw("select image from inventories where category_id = $1 limit 2", CategoryID).Scan(&images).Error
	if err != nil {
		return []string{}, err
	}

	return images, nil

}
