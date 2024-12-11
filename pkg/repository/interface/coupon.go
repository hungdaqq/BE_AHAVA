package interfaces

import (
	"ahava/pkg/domain"
	"ahava/pkg/utils/models"
)

type CouponRepository interface {
	AddCoupon(models.Coupons) error
	MakeCouponInvalid(id int) error
	ReActivateCoupon(id int) error
	FindCouponDetails(couponID int) (domain.Coupons, error)
	GetAllCoupons() ([]domain.Coupons, error)
}
