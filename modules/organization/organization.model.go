package organization

import (
	"myapp/database"
	"myapp/modules/user"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Address string      `json:"address"`
	Social  string      `json:"social"`
	Users   []user.User `json:"users" gorm:"-"`
}

func OrganizationList() []Organization {
	db := database.DB

	orgs := []Organization{}

	// users := make([]user.User, 0)
	for i := range orgs {
		users := make([]user.User, 0)
		db.Table("users").Joins("INNER JOIN userorgs on userorgs.user_id = users.id").Select("users.*").Find(&users)

		orgs[i].Users = users
	}
	// orgs.Users = users
	// copy(orgs.Users,users)
	return orgs
}

func FindOrganization(org_id uint) Organization {
	db := database.DB

	org := Organization{}

	db.Where("id = ?", org_id).First(&org)

	users := make([]user.User, 0)
	db.Table("users").Joins("INNER JOIN userorgs on userorgs.user_id = users.id").Where("userorgs.organization_id = ?", org_id).Select("users.*").Find(&users)
	org.Users = users

	return org
}
