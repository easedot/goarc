package presenter

//定义表现层抽象接口
import (
	"github.com/easedot/hb_vendor/domain"
)

type VendorPresenter interface {
	ResponseVendors(as []*domain.Vendor) []*domain.Vendor
	ResponseVendor(as *domain.Vendor) *domain.Vendor
}
type UserPresenter interface {
	ResponseUsers(as []*domain.User) []*domain.User
	ResponseUser(as *domain.User) *domain.User
}
