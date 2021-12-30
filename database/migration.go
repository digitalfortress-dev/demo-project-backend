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
			// Rollback: func(tx *gorm.DB) error {
			// 	if err := tx.DropTable("model").Error; err != nil {
			// 		return nil
			// 	}
			// 	return nil
			// },
		},
	})

}
