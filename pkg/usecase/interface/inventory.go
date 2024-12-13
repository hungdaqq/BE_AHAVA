package interfaces

import (
	"ahava/pkg/utils/models"
	"mime/multipart"
)

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventory, image *multipart.FileHeader) (models.InventoryResponse, error)
	UpdateInventory(int, models.UpdateInventory) error
	DeleteInventory(id string) error

	ShowIndividualProducts(sku string) (models.Inventories, error)
	ListProductsForUser(page, userID int) ([]models.Inventories, error)
	ListProductsForAdmin(page int) ([]models.Inventories, error)

	SearchProducts(key string) ([]models.Inventories, error)
	UpdateProductImage(id int, file *multipart.FileHeader) error
}
