package models

type InventoryResponse struct {
	ID          int
	ProductName string
	Stock       int
}

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}

type Inventories struct {
	ID                  uint    `json:"id"`
	CategoryID          int     `json:"category_id"`
	Image               string  `json:"image"`
	ProductName         string  `json:"product_name"`
	Size                string  `json:"size"`
	Stock               int     `json:"stock"`
	Price               float64 `json:"price"`
	IfPresentAtWishlist bool    `json:"if_present_at_wishlist"`
	IfPresentAtCart     bool    `json:"if_present_at_cart"`
	DiscountedPrice     float64 `json:"discounted_price"`
}

type AddInventory struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type UpdateInventory struct {
	ID         int     `json:"id"`
	Name       string  `json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	Stock      int     `json:"stock" validate:"required"`
	CategoryID int     `json:"category_id" validate:"required"`
	Size       string  `json:"size"`
}

type Banner struct {
	CategoryID         int
	CategoryName       string
	DiscountPercentage int
	Images             []string `gorm:"-"`
}
