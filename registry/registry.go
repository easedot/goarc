package registry

import (
	"github.com/easedot/godbs"

	"github.com/easedot/hb_vendor/adapters/controller"
)

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *godbs.DbHelper) Registry {
	return &registry{db: db}
}

type registry struct {
	db *godbs.DbHelper
}

func (r *registry) NewAppController() controller.AppController {
	c := controller.Controller{
		VendorController: r.NewVendorController(),
		UserController:   r.NewUserController()}
	return c
}
