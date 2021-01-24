package presenter

import (
	"github.com/easedot/hb_vendor/domain"
	"github.com/easedot/hb_vendor/usecases/presenter"
)

//实现存的具体实现，通过抽象接口返回，使用层就可以基于接口，不依赖具体实现
func NewVendorPresenter() presenter.VendorPresenter {
	return &vendorPresenter{}
}

type vendorPresenter struct {
}

func (ap *vendorPresenter) ResponseVendors(as []*domain.Vendor) []*domain.Vendor {
	//for _, a := range as {
	//a.Name = "EASE_" + a.Name
	//}
	return as
}
func (ap *vendorPresenter) ResponseVendor(as *domain.Vendor) *domain.Vendor {
	//as.Name = "EASE_" + as.Name
	return as
}

//实现存的具体实现，通过抽象接口返回，使用层就可以基于接口，不依赖具体实现
func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

type userPresenter struct {
}

func (ap *userPresenter) ResponseUsers(as []*domain.User) []*domain.User {
	//for _, a := range as {
	//	a.Name = a.Name
	//}
	return as
}
func (ap *userPresenter) ResponseUser(as *domain.User) *domain.User {
	//as.Name =  as.Name
	return as
}
