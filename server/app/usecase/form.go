package usecase

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/model"
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
)

type FormUsecase interface {
	GetAllPaymentTypes() ([]model.PaymentType, error)
	GetAllPaymentStatuses() ([]model.PaymentStatus, error)
	GetAllMonths() ([]model.Month, error)
}

func NewFormUsecase(
	paymentTypeRepo repository.PaymentTypeRepository,
	paymentStatusRepo repository.PaymentStatusRepository,
	monthRepo repository.MonthRepository) FormUsecase {

	return implFormUsecase{
		PaymentTypeRepository:   paymentTypeRepo,
		PaymentStatusRepository: paymentStatusRepo,
		MonthRepository:         monthRepo,
	}
}

type implFormUsecase struct {
	PaymentTypeRepository   repository.PaymentTypeRepository
	PaymentStatusRepository repository.PaymentStatusRepository
	MonthRepository         repository.MonthRepository
}

func (self implFormUsecase) GetAllPaymentTypes() ([]model.PaymentType, error) {
	ptl, err := self.PaymentTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return ptl, nil
}

func (self implFormUsecase) GetAllPaymentStatuses() ([]model.PaymentStatus, error) {
	psl, err := self.PaymentStatusRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return psl, nil
}

func (self implFormUsecase) GetAllMonths() ([]model.Month, error) {
	ml, err := self.MonthRepository.GetAll()
	if err != nil {
		return nil, err

	}

	return ml, nil
}
