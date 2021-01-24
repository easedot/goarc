package controller

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/easedot/hb_vendor/domain"
	"github.com/easedot/hb_vendor/usecases/interactor"
	"github.com/easedot/hb_vendor/utils"
)

type UserController interface {
	ResetPassword(c Context) error
	ChangePassword(c Context) error
	DisableUser(c Context) error
	EnableUser(c Context) error
	GetUsers(c Context) error
	UpdateUser(c Context) error
	CreateUser(c Context) error
	SignIn(c Context) error
	SignOut(c Context) error
	SignUp(c Context) error
	CaptchaGen(c Context) error
	CaptchaImg(c Context) error
}

type userController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(ai interactor.UserInteractor) UserController {
	return &userController{ai}
}
func (uc *userController) CaptchaGen(c Context) error {
	id := captcha.NewLen(4)
	return c.JSON(http.StatusOK, map[string]string{
		"id": id,
	})
}
func (uc *userController) CaptchaImg(c Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.ErrBadRequest
	}
	var imageBuf bytes.Buffer
	captcha.Reload(id)
	captcha.WriteImage(&imageBuf, id, 110, 50)
	return c.Stream(http.StatusOK, "image/png", &imageBuf)
}
func (uc *userController) SignIn(c Context) error {
	// Throws unauthorized error
	param := &domain.User{}
	if err := c.Bind(param); err != nil {
		return err
	}
	param.Offset = 0
	param.Limit = 1
	if !captcha.VerifyString(param.CaptchaId, param.CaptchaCode) {
		return echo.ErrUnauthorized
	}
	password := param.Password
	param.Password = ""
	users, err := uc.userInteractor.Query(param)
	if err != nil {
		return err
	}
	if len(users) <= 0 {
		return echo.ErrUnauthorized
	}
	user := users[0]
	if user.State != domain.UserApprove {
		return echo.ErrUnauthorized
	}
	if !utils.ComparePasswords(user.Password, password) {
		return echo.ErrUnauthorized
	}

	t, err := uc.genToken(user)
	if err != nil {
		return err
	}
	roles := []string{"ROLE_VENDOR"}
	if user.Type == domain.ADMIN {
		roles = []string{"ROLE_ADMIN"}
	}
	if user.Type == domain.OPER {
		roles = []string{"ROLE_OPER"}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": t,
		"name":        user.Name,
		"type":        user.Type,
		"create_time": user.CreateTime,
		"state":       user.State,
		"roles":       roles,
	})
}

func (uc *userController) genToken(user *domain.User) (string, error) {
	// Set custom claims
	claims := &domain.JwtCustomClaims{
		user.ID,
		user.Type,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("hb_vendor_secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (uc *userController) SignOut(c Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"ok": "ok",
	})
}
func (uc *userController) SignUp(c Context) error {
	u := &domain.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	if !captcha.VerifyString(u.CaptchaId, u.CaptchaCode) {
		return captcha.ErrNotFound
	}

	u.Password = utils.HashAndSalt(u.Password)
	u.Type = domain.NORMAL
	u.State = domain.UserApprove
	if err := uc.userInteractor.Create(u); err != nil {
		return echo.ErrInternalServerError
	}

	t, err := uc.genToken(u)
	if err != nil {
		return err
	}
	roles := []string{"ROLE_VENDOR"}
	if u.Type == domain.ADMIN {
		roles = []string{"ROLE_ADMIN"}
	}
	if u.Type == domain.OPER {
		roles = []string{"ROLE_OPER"}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": t,
		"name":        u.Name,
		"type":        u.Type,
		"create_time": u.CreateTime,
		"state":       u.State,
		"roles":       roles,
	})
}

func (uc *userController) GetUsers(c Context) error {
	type User struct {
		Page    int    `json:"page" form:"page" query:"page"`
		OrderBy string `json:"order_by" form:"order_by" query:"order_by"`
	}
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	if claims.Type == domain.NORMAL {
		return echo.ErrForbidden
	}
	//default offset 0 limit 10
	ar := &domain.User{Offset: u.Page * 5, Limit: 5, OrderBy: u.OrderBy}
	if claims.Type == domain.ADMIN {
		ar.Type = domain.OPER
	}
	if claims.Type == domain.OPER {
		ar.Type = domain.NORMAL
	}
	us, err := uc.userInteractor.RawQuery(ar)
	if err != nil {
		return err
	}
	for _, u := range us {
		u.Password = ""
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": us,
		"meta":  map[string]interface{}{"offset": ar.Offset, "limit": ar.Limit},
	})
}

func (uc *userController) ChangePassword(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	type User struct {
		OldPassword string `json:"old_password" form:"old_password" query:"old_password"`
		NewPassword string `json:"new_password" form:"new_password" query:"new_password"`
	}
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	ar := &domain.User{ID: claims.UserId}

	fu, err := uc.userInteractor.Find(ar)
	if err != nil {
		return err
	}
	if !utils.ComparePasswords(fu.Password, u.OldPassword) {
		return echo.ErrUnauthorized
	}

	ar.Password = utils.HashAndSalt(u.NewPassword)
	err = uc.userInteractor.UpdatePassword(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)

}

func (uc *userController) ResetPassword(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	if claims.Type != domain.ADMIN && claims.Type != domain.OPER {
		return echo.ErrForbidden
	}

	id := c.Param("id")
	idi, _ := strconv.ParseInt(id, 10, 64)
	passDefault := "Qazwsx123"
	ar := &domain.User{ID: idi}
	ar.Password = utils.HashAndSalt(passDefault)
	err := uc.userInteractor.UpdatePassword(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)
}

func (uc *userController) EnableUser(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	if claims.Type != domain.ADMIN && claims.Type != domain.OPER {
		return echo.ErrForbidden
	}

	id := c.Param("id")
	idi, _ := strconv.ParseInt(id, 10, 64)
	ar := &domain.User{ID: idi, State: domain.UserApprove}
	err := uc.userInteractor.UpdateState(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)

}
func (uc *userController) DisableUser(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	if claims.Type != domain.ADMIN && claims.Type != domain.OPER {
		return echo.ErrForbidden
	}

	id := c.Param("id")
	idi, _ := strconv.ParseInt(id, 10, 64)
	ar := &domain.User{ID: idi, State: domain.UserReject}
	err := uc.userInteractor.UpdateState(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)
}

func (uc *userController) UpdateUser(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	if claims.Type != domain.ADMIN {
		return echo.ErrForbidden
	}
	ar := &domain.User{}
	c.Bind(ar)
	err := uc.userInteractor.Update(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)
}

func (uc *userController) CreateUser(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	if claims.Type != domain.ADMIN {
		return echo.ErrForbidden
	}
	ar := &domain.User{}
	c.Bind(ar)
	ar.Type = domain.OPER
	ar.State = domain.UserApprove
	ar.Password = utils.HashAndSalt(ar.Password)
	err := uc.userInteractor.Create(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)
}
