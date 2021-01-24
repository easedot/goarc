package interactor

import (
	"github.com/easedot/hb_vendor/domain"
	"github.com/easedot/hb_vendor/usecases/presenter"
	"github.com/easedot/hb_vendor/usecases/repository"
)

//业务逻辑层,这里面对的都是接口，具体实现延后
//for use case business logic
type VendorInteractor interface {
	Query(u []*domain.Vendor) ([]*domain.Vendor, error)
	Find(u *domain.Vendor) (*domain.Vendor, error)
	Update(u *domain.Vendor) error
	UpdateState(u *domain.Vendor) error
	Create(u *domain.Vendor) error
}

//inject from outer object to interface
func NewVendorInteractor(r repository.VendorRepository, p presenter.VendorPresenter) VendorInteractor {
	return &vendorInteractor{p, r}
}

type vendorInteractor struct {
	VendorPresenter  presenter.VendorPresenter
	VendorRepository repository.VendorRepository
}

func (ai *vendorInteractor) Query(u []*domain.Vendor) ([]*domain.Vendor, error) {
	a, err := ai.VendorRepository.Query(u)
	if err != nil {
		return nil, err
	}
	//user case modify response,use inject adapter response object
	a = ai.VendorPresenter.ResponseVendors(a)
	return a, nil
}

func (ai *vendorInteractor) Find(u *domain.Vendor) (*domain.Vendor, error) {
	a, err := ai.VendorRepository.Find(u)
	if err != nil {
		return nil, err
	}
	//user case modify response,use inject adapter response object
	a = ai.VendorPresenter.ResponseVendor(a)
	return a, nil
}
func (ai *vendorInteractor) Update(u *domain.Vendor) error {
	err := ai.VendorRepository.Update(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	u = ai.VendorPresenter.ResponseVendor(u)
	return nil
}
func (ai *vendorInteractor) UpdateState(u *domain.Vendor) error {
	err := ai.VendorRepository.UpdateState(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	u = ai.VendorPresenter.ResponseVendor(u)
	return nil
}
func (ai *vendorInteractor) Create(u *domain.Vendor) error {
	err := ai.VendorRepository.Create(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	u = ai.VendorPresenter.ResponseVendor(u)
	return nil
}
