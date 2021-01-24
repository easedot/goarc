package controller

type AppController interface {
	VendorController
	UserController
}

type Controller struct {
	VendorController
	UserController
}
