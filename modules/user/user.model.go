package user

import (
	"myapp/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	// Vehicle []vehicle.Vehicle `json:"vehicle" gorm:"foreignKey:VehicleId"`
}

func (u *User) Create() error {
	db := database.DB
	return db.Create(&u).Error
}

func (u *User) Update() error {
	db := database.DB
	return db.Updates(&u).Error
}

func (u *User) Delete() error {
	db := database.DB
	return db.Delete(&u).Error
}
