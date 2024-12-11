package interfaces

import (
	"ahava/pkg/domain"
	"ahava/pkg/utils/models"
)

type OfferUseCase interface {
	AddNewOffer(model models.OfferMaking) error
	MakeOfferExpire(id int) error
	GetOffers() ([]domain.Offer, error)
}
