package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/savsgio/go-logger"

	"github.com/easedot/hb_vendor/domain"
	"github.com/easedot/hb_vendor/usecases/interactor"
)

type VendorController interface {
	GetVendors(c Context) error
	GetVendor(c Context) error
	UpdateVendor(c Context) error
	CreateVendor(c Context) error
	UploadFile(c Context) error
	UploadFiles(c Context) error
	DeleteUploadFiles(c Context) error
	GetUploadFiles(c Context) error
}

type vendorController struct {
	//call usecase to implete logic
	vendorInteractor interactor.VendorInteractor
}

func NewVendorController(ai interactor.VendorInteractor) VendorController {
	return &vendorController{ai}
}
func (ac *vendorController) UploadFile(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename := fmt.Sprintf("%d/success_case_%s", claims.UserId, file.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"files": filename})

}
func (ac *vendorController) GetUploadFiles(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	filename := c.Param("id")
	dir := fmt.Sprintf("uplodad/%d/", claims.UserId)
	newFile := filepath.Join(dir, filepath.Base(filename))
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		logger.Info("file not exists %s", newFile)
		return echo.ErrNotFound
	} else {
		return c.File(newFile)
	}
}

func (ac *vendorController) DeleteUploadFiles(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	filename := c.Param("id")
	dir := fmt.Sprintf("uplodad/%d/", claims.UserId)
	newFile := filepath.Join(dir, filepath.Base(filename))
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		logger.Info("file not exists %s", newFile)
	} else {
		var err = os.Remove(newFile)
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, "ok")
}

func (ac *vendorController) UploadFiles(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	name := c.FormValue("name")

	files := form.File["files[]"]
	var filenames []string
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dir := fmt.Sprintf("uplodad/%d/", claims.UserId)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
		filename := fmt.Sprintf("%s_%s", name, file.Filename)
		newFile := filepath.Join(dir, filepath.Base(filename))
		dst, err := os.Create(newFile)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		filenames = append(filenames, filename)
	}
	return c.JSON(http.StatusOK, filenames)
}

func (ac *vendorController) GetVendors(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)

	var ars []*domain.Vendor
	ar := &domain.Vendor{UserId: claims.UserId}
	ars = append(ars, ar)
	a, err := ac.vendorInteractor.Query(ars)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func (ac *vendorController) GetVendor(c Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ar := &domain.Vendor{ID: int64(id)}
	a, err := ac.vendorInteractor.Find(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func (ac *vendorController) CreateVendor(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	ar := &domain.Vendor{}
	c.Bind(ar)
	ar.UserId = claims.UserId
	err := ac.vendorInteractor.Create(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)
}

func (ac *vendorController) UpdateVendor(c Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	ar := &domain.Vendor{}
	c.Bind(ar)
	ar.UserId = claims.UserId
	err := ac.vendorInteractor.Update(ar)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ar)
}
