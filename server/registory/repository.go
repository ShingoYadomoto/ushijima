package registory

import (
	"github.com/ShingoYadomoto/ushijima/server/domain/repository"
	"github.com/ShingoYadomoto/ushijima/server/infrastructure/datastore/impl_repository"
)

// repository
func (self register) NewMonthRepository() repository.MonthRepository {
	return impl_repository.NewMonthRepository(self.db)
}

func (self register) NewPaymentRepository() repository.PaymentRepository {
	return impl_repository.NewPaymentRepository(self.db)
}

func (self register) NewPaymentStatusRepository() repository.PaymentStatusRepository {
	return impl_repository.NewPaymentStatusRepository(self.db)
}

func (self register) NewPaymentTypeRepository() repository.PaymentTypeRepository {
	return impl_repository.NewPaymentTypeRepository(self.db)
}
