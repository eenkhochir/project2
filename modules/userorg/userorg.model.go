package userorg

import (
	"gorm.io/gorm"
)

type Userorg struct {
	gorm.Model
	UserID         int
	OrganizationID int
}
