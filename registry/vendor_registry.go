package registry

import (
	"github.com/easedot/hb_vendor/adapters/controller"
	"github.com/easedot/hb_vendor/adapters/presenter"
	"github.com/easedot/hb_vendor/adapters/repository"
	"github.com/easedot/hb_vendor/usecases/interactor"
)

func (r *registry) NewVendorController() controller.VendorController {
	ar := repository.NewVendorRepository(r.db)
	ap := presenter.NewVendorPresenter()
	ai := interactor.NewVendorInteractor(ar, ap)
	return controller.NewVendorController(ai)
}
func (r *registry) NewUserController() controller.UserController {
	ar := repository.NewUserRepository(r.db)
	ap := presenter.NewUserPresenter()
	ai := interactor.NewUserInteractor(ar, ap)
	return controller.NewUserController(ai)
}
