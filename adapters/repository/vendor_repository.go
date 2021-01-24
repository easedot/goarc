package repository

//这里是存储的具体实现，落实的具体的存储类型，如Mysql、Oracle
import (
	"fmt"
	"strconv"
	"strings"

	"github.com/easedot/godbs"
	"github.com/savsgio/go-logger"

	"github.com/easedot/hb_vendor/domain"
	"github.com/easedot/hb_vendor/usecases/repository"
)

type vendorRepository struct {
	db *godbs.DbHelper //存储实现
}

//注入存储实现，返回存储接口协议
func NewVendorRepository(db *godbs.DbHelper) repository.VendorRepository {
	return &vendorRepository{db: db} //结构体要实现这些接口才能作为接口返回
}

//实现存储接口的定义
func (ar *vendorRepository) Query(u []*domain.Vendor) ([]*domain.Vendor, error) {
	var rst []*domain.Vendor
	if err := ar.db.Query(u[0], &rst); err != nil {
		return nil, err
	}
	return rst, nil
}
func (ar *vendorRepository) Find(u *domain.Vendor) (*domain.Vendor, error) {
	if err := ar.db.Find(u); err != nil {
		return nil, err
	}
	return u, nil
}
func (ar *vendorRepository) Update(u *domain.Vendor) error {
	qv := fmt.Sprintf("select id from vendor where user_id= %d", u.UserId)
	vl, err := ar.db.SqlMap(qv)
	if err != nil {
		return err
	}
	if vl == nil {
		return domain.ErrNotFound
	}
	logger.Info("find vendor %d", len(vl))
	u.ID, _ = strconv.ParseInt(vl[0]["id"], 0, 64)
	//u.CreateTime = sql.NullTime{}
	//u.UpdateTime = sql.NullTime{}
	if err := ar.db.Update(u); err != nil {
		return err
	}
	return nil
}
func (ar *vendorRepository) Create(u *domain.Vendor) error {
	if err := ar.db.Create(u); err != nil {
		return err
	}
	return nil
}

type userRepository struct {
	db *godbs.DbHelper //存储实现
}

//注入存储实现，返回存储接口协议
func NewUserRepository(db *godbs.DbHelper) repository.UserRepository {
	return &userRepository{db: db} //结构体要实现这些接口才能作为接口返回
}

//实现存储接口的定义
func (ar *userRepository) RawQuery(u *domain.User) ([]*domain.User, error) {
	_, _, _, w := ar.db.Info(u)
	wh := strings.Join(w, " and ")
	var r []*domain.User
	q := fmt.Sprintf("where %s limit %d, %d ", wh, u.Offset, u.Limit)
	if err := ar.db.SqlStructSlice(q, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (ar *userRepository) Query(u *domain.User) ([]*domain.User, error) {
	var r []*domain.User
	if err := ar.db.Query(u, &r); err != nil {
		return nil, err
	}
	//var r []*domain.User
	//q := fmt.Sprintf("where type=%d   limit %d, %d ", u.Type, u.Offset, u.Limit)
	//if err := ar.db.SqlStructSlice(q, &r); err != nil {
	//	return nil, err
	//}
	return r, nil
}
func (ar *userRepository) Find(u *domain.User) (*domain.User, error) {
	if err := ar.db.Find(u); err != nil {
		return nil, err
	}
	return u, nil
}
func (ar *userRepository) Update(u *domain.User) error {
	if err := ar.db.Update(u); err != nil {
		return err
	}
	return nil
}
func (ar *userRepository) Create(u *domain.User) error {
	if err := ar.db.Create(u); err != nil {
		return err
	}
	return nil
}
