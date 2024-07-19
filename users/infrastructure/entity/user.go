package entity

import "clockify/project/domain"

type User struct {
	ID       int64  `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"type:varchar(255);size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	//Password string            `gorm:"size:255;not null;" json:"password;column:password"`
	projects []domain.Projects `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
