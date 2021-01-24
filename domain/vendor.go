package domain

import (
	"database/sql"
)

// Article represent the article model
type Vendor struct {
	ID                     int64        `json:"id"`
	UserId                 int64        `json:"user_id"`
	State                  int          `json:"state"`
	Type                   int          `json:"type"`
	Name                   string       `json:"name"`
	Address                string       `json:"address"`
	Phone                  string       `json:"phone"`
	CountryCode            string       `json:"country_code"`
	WebLink                string       `json:"web_link"`
	RegisteredAddress      string       `json:"registered_address"`
	RegisteredCapital      int          `json:"registered_capital"`
	RegisteredNo           string       `json:"registered_no"`
	RegisteredDate         string       `json:"registered_date"`
	RegisteredType         int          `json:"registered_type"`
	TaxNo                  string       `json:"tax_no"`
	EmployeeCount          int          `json:"employee_count"`
	MarketStaffCount       int          `json:"market_staff_count"`
	TechnicalStaffCount    int          `json:"technical_staff_count"`
	BankName               string       `json:"bank_name"`
	BankAccount            string       `json:"bank_account"`
	Referrer               string       `json:"referrer"`
	ReferrerReason         string       `json:"referrer_reason"`
	SuccessCaseDocuments   string       `json:"success_case_documents"`
	MainProduct            string       `json:"main_product"`
	ChannelLevel           string       `json:"channel_level"`
	IsAllCountry           int          `json:"is_all_country"`
	BossName               string       `json:"boss_name"`
	BossEmail              string       `json:"boss_email"`
	BossPhone              string       `json:"boss_phone"`
	BossTel                string       `json:"boss_tel"`
	ContactName            string       `json:"contact_name"`
	ContactEmail           string       `json:"contact_email"`
	ContactPhone           string       `json:"contact_phone"`
	ContactTel             string       `json:"contact_tel"`
	QualificationDocuments string       `json:"qualification_documents"`
	CreateTime             sql.NullTime `json:"create_time" db:"r"`
	UpdateTime             sql.NullTime `json:"update_time" db:"r"`
}
