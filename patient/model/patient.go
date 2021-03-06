package patient

type (
	Patient struct {
		FirstName         string `gorm:"not null" json:"first_name" validate:"required"`
		LastNanme         string `gorm:"not null" json:"last_name" validate:"required"`
		DateOfBirth       int    `json:"date_of_birth"`
		PhoneNumber       string `json:"phone_number" validate:"required"`
		Email             string `json:"email"  validate:"required"`
		Address           string `json:"address"`
		AppointmentOfTime int    `json:"appointment_time"`
		SrcPicture        string `json:"src_picture"`
		MyId              int    `gorm:"primary_key;auto_increment;not_null" json:"id"`
	}
)
