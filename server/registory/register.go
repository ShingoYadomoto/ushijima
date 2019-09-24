package registory

import (
	"github.com/ShingoYadomoto/ushijima/server/app/handler"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/postgres"
)

type Register interface {
	NewHomeHandler() handler.HomeHandler
	NewPaymentHandler() handler.PaymentHandler
	NewAppHandler() handler.AppHandler
}

func NewRegister(db postgres.AbsDB) Register {
	return &register{db}
}

type register struct {
	db postgres.AbsDB
}

type appHandler struct {
	handler.HomeHandler
	handler.FormHandler
	handler.PaymentHandler
}

func (self register) NewAppHandler() handler.AppHandler {
	return &appHandler{
		HomeHandler:    self.NewHomeHandler(),
		FormHandler:    self.NewFormHandler(),
		PaymentHandler: self.NewPaymentHandler(),
	}
}
