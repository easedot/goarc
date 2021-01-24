package interactor

import (
	"github.com/easedot/hb_vendor/domain"
	"github.com/easedot/hb_vendor/usecases/presenter"
	"github.com/easedot/hb_vendor/usecases/repository"
)

//用户的交互逻辑
type UserInteractor interface {
	RawQuery(u *domain.User) ([]*domain.User, error)
	Query(u *domain.User) ([]*domain.User, error)
	Find(u *domain.User) (*domain.User, error)
	Update(u *domain.User) error
	Create(u *domain.User) error
	UpdateState(u *domain.User) error
	UpdatePassword(u *domain.User) error
}

//inject from outer object to interface
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{p, r}
}

type userInteractor struct {
	UserPresenter  presenter.UserPresenter
	UserRepository repository.UserRepository
}

func (ai *userInteractor) RawQuery(u *domain.User) ([]*domain.User, error) {
	a, err := ai.UserRepository.RawQuery(u)
	if err != nil {
		return nil, err
	}
	//user case modify response,use inject adapter response object
	a = ai.UserPresenter.ResponseUsers(a)
	return a, nil
}

func (ai *userInteractor) Query(u *domain.User) ([]*domain.User, error) {
	a, err := ai.UserRepository.Query(u)
	if err != nil {
		return nil, err
	}
	//user case modify response,use inject adapter response object
	a = ai.UserPresenter.ResponseUsers(a)
	return a, nil
}

func (ai *userInteractor) Find(u *domain.User) (*domain.User, error) {
	a, err := ai.UserRepository.Find(u)
	if err != nil {
		return nil, err
	}
	//user case modify response,use inject adapter response object
	a = ai.UserPresenter.ResponseUser(a)
	return a, nil
}
func (ai *userInteractor) Update(u *domain.User) error {
	err := ai.UserRepository.Update(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	//u = ai.UserPresenter.ResponseUser(u)
	return nil
}

func (ai *userInteractor) Create(u *domain.User) error {
	err := ai.UserRepository.Create(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	u = ai.UserPresenter.ResponseUser(u)
	return nil
}
func (ai *userInteractor) UpdateState(u *domain.User) error {
	err := ai.UserRepository.UpdateState(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	//u = ai.UserPresenter.ResponseUser(u)
	return nil
}
func (ai *userInteractor) UpdatePassword(u *domain.User) error {
	err := ai.UserRepository.UpdatePassword(u)
	if err != nil {
		return err
	}
	//user case modify response,use inject adapter response object
	//u = ai.UserPresenter.ResponseUser(u)
	return nil
}
