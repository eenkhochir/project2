package vehicle

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	Brand        string `json:"name"`
	Year         int    `json:"year"`
	UserId       uint   `json:"user_id"`
	Owner        string `json:"owner"`
	LicensePlate string `json:"licensePlate"`
}
