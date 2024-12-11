package interfaces

import (
	"ahava/pkg/domain"
	"ahava/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(model models.OfferMaking) error
	MakeOfferExpire(id int) error
	FindDiscountPercentage(int) (int, error)
	GetOffers() ([]domain.Offer, error)
}
