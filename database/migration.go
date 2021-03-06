package database

import (
	Patient "main/patient/model"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {

	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20212412",
			Migrate: func(transmit *gorm.DB) error {
				if err := transmit.AutoMigrate(&Patient.Patient{}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	})

}

func CheckTableInDatabase(db *gorm.DB) {

	db.DropTable("migrations")
	db.DropTableIfExists(&Patient.Patient{})
}
