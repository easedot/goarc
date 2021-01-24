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
	UserApprove int = iota + 1
	UserReject
)

type VendorState int

///*1:待审核 2:审核中 3:通过 4:拒绝 5:补充*/
const (
	VendorPending int = iota + 1
	VendorSubmited
	VendorApproveed
	VendorReject
	VendorReWrite
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
