package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	UserId int64 `json:"user_id"`
	Type   int   `json:"type"`
	jwt.StandardClaims
}

type UserState int

const (
	UserPending int = iota + 1
	UserApprove
	UserReject
)

type VendorState int

const (
	VendorPending UserState = iota + 1
	VendorApprove
	VendorReject
)

type VendorType int

const (
	VendorNormal UserState = iota + 1
	VendorTemp
)

const (
	NORMAL int = iota + 1
	ADMIN
	OPER
)
