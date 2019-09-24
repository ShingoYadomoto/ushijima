package registory

import "github.com/ShingoYadomoto/ushijima/server/app/handler"

// handler
func (self register) NewHomeHandler() handler.HomeHandler {
	return handler.NewHomeHandler()
}

func (self register) NewFormHandler() handler.FormHandler {
	return handler.NewFormHandler(self.NewFormUsecase())
}

func (self register) NewPaymentHandler() handler.PaymentHandler {
	return handler.NewPaymentHandler(self.NewPaymentUsecase())
}
