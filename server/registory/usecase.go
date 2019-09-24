package registory

import "github.com/ShingoYadomoto/ushijima/server/app/usecase"

// usecase
func (self register) NewFormUsecase() usecase.FormUsecase {
	return usecase.NewFormUsecase(
		self.NewPaymentTypeRepository(),
		self.NewPaymentStatusRepository(),
		self.NewMonthRepository(),
	)
}

func (self register) NewPaymentUsecase() usecase.PaymentUsecase {
	return usecase.NewPaymentUsecase(
		self.NewPaymentRepository(),
		self.NewMonthRepository(),
	)
}
