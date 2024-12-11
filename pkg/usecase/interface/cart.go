package interfaces

import "ahava/pkg/utils/models"

type CartUseCase interface {
	AddToCart(user_id, inventory_id int) error
	CheckOut(id int) (models.CheckOut, error)
}
