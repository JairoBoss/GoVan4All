package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Driver struct {
	ID        string    `gorm:"primary_key" json:"id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	CompanyID string    `gorm:"not null" json:"company_id"`
	User      User      `gorm:"foreignkey:UserID" json:"user"`
	Company   Company   `gorm:"foreignkey:CompanyID" json:"company"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (d *Driver) SaveDriver(db *gorm.DB) (*Driver, error) {
	var err error
	err = db.Debug().Create(&d).Error
	if err != nil {
		return &Driver{}, err
	}
	return d, nil
}

func (d *Driver) FindAllDrivers(db *gorm.DB) (*[]Driver, error) {
	var err error
	drivers := []Driver{}
	err = db.Debug().Model(&Driver{}).Find(&drivers).Error
	if err != nil {
		return &[]Driver{}, err
	}
	return &drivers, nil
}

func (d *Driver) FindDriverByID(db *gorm.DB, id string) (*Driver, error) {
	var err error
	err = db.Debug().Model(Driver{}).Where("id = ?", id).Take(&d).Error
	if err != nil {
		return &Driver{}, err
	}
	return d, nil
}

func (d *Driver) UpdateDriver(db *gorm.DB) (*Driver, error) {
	err := db.Debug().Model(&Driver{}).Where("id = ?", d.ID).Take(&Driver{}).Updates(
		Driver{
			UserID:    d.UserID,
			CompanyID: d.CompanyID,
			UpdatedAt: time.Now(),
		},
	).Error
	if err != nil {
		return &Driver{}, err
	}
	return d, nil
}

func (d *Driver) DeleteDriver(db *gorm.DB, id string) error {
	err := db.Debug().Model(&Driver{}).Where("id = ?", id).Take(&Driver{}).Delete(&Driver{}).Error
	if err != nil {
		return err
	}
	return nil
}
