package interfaces

import "ahava/pkg/utils/models"

type OtpUseCase interface {
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
	SendOTP(phone string) error
}
