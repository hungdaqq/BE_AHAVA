package interfaces

import (
	"ahava/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.AddInventory, url string) (models.InventoryResponse, error)
	CheckInventory(pid int) (bool, error)
	UpdateInventory(int, models.UpdateInventory) error
	DeleteInventory(id string) error

	ShowIndividualProducts(id string) (models.Inventories, error)
	ListProducts(page int) ([]models.Inventories, error)
	ListProductsByCategory(id int) ([]models.Inventories, error)
	CheckStock(inventory_id int) (int, error)
	CheckPrice(inventory_id int) (float64, error)
	SearchProducts(key string) ([]models.Inventories, error)
	UpdateProductImage(int, string) error
}
