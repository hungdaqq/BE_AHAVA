package interfaces

import (
	"ahava/pkg/domain"
	"ahava/pkg/utils/models"
)

type CouponUsecase interface {
	AddCoupon(coupon models.Coupons) error
	MakeCouponInvalid(id int) error
	ReActivateCoupon(id int) error
	GetAllCoupons() ([]domain.Coupons, error)
}
