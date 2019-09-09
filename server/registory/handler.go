package registory

import "github.com/ShingoYadomoto/ushijima/server/app/handler"

// handler
func (self register) NewHomeHandler() handler.HomeHandler {
	return handler.NewHomeHandler(
		self.NewPaymentTypeRepository(),
		self.NewPaymentStatusRepository(),
		self.NewMonthRepository(),
	)
}

func (self register) NewPaymentHandler() handler.PaymentHandler {
	return handler.NewPaymentHandler(
		self.NewPaymentRepository(),
		self.NewMonthRepository(),
	)
}
