package interfaces

import "ahava/pkg/utils/models"

type WishlistUseCase interface {
	AddToWishlist(userID, InventoryID int) error
	RemoveFromWishlist(invID, userID int) error
	GetWishList(id int) ([]models.Inventories, error)
}
