package domain

// User represent the author model
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	State       int    `json:"state"`
	Type        int    `json:"type"`
	Password    string `json:"password"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	CaptchaId   string `json:"captcha_id" db:"-"`
	CaptchaCode string `json:"captcha_code" db:"-"`
	Offset      int    `json:"offset" db:"-"`
	Limit       int    `json:"limit" db:"-"`
}
