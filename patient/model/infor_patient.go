package model

import "github.com/jinzhu/gorm"

type (
	Patient struct {
		FirstName         string `gorm:"not null" json:"first_name" validate:"required"`
		LastNanme         string `gorm:"not null" json:"last_name" validate:"required"`
		DateOfBirth       string    `json:"date_of_birth"`
		PhoneNumber       string    `json:"phone_number" validate:"required"`
		Email             string `json:"email"  validate:"required"`
		Address           string `json:"address"`
		AppointmentOfTime int    `json:"appointment_time"`
		SrcPicture        string `json:"src_picture"`
		gorm.Model
	}
)

type (
	User struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		AccessToken string `json:"access_token"`
	}
)
