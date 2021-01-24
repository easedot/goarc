package repository

//定义业务抽象存储接口协议
import (
	"github.com/easedot/hb_vendor/domain"
)

type VendorRepository interface {
	Query(u []*domain.Vendor) ([]*domain.Vendor, error)
	Find(u *domain.Vendor) (*domain.Vendor, error)
	Update(u *domain.Vendor) error
	Create(u *domain.Vendor) error
}

type UserRepository interface {
	RawQuery(u *domain.User) ([]*domain.User, error)
	Query(u *domain.User) ([]*domain.User, error)
	Find(u *domain.User) (*domain.User, error)
	Update(u *domain.User) error
	Create(u *domain.User) error
}
