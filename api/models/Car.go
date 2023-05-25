package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Car struct {
	ID            string    `gorm:"primary_key" json:"id"`
	Model         string    `gorm:"not null" json:"model"`
	Brand         string    `gorm:"not null" json:"brand"`
	Year          int       `gorm:"not null" json:"year"`
	NumberOfDoors int       `gorm:"not null" json:"number_of_doors"`
	NumberOfSeats int       `gorm:"not null" json:"number_of_seats"`
	CompanyID     string    `gorm:"not null" json:"company_id"`
	DriverID      string    `gorm:"not null" json:"driver_id"`
	Company       Company   `gorm:"foreignkey:CompanyID" json:"company"`
	Driver        Driver    `gorm:"foreignkey:DriverID" json:"driver"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Car) SaveCar(db *gorm.DB) (*Car, error) {
	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Car{}, err
	}
	return c, nil
}

func (c *Car) FindAllCars(db *gorm.DB) (*[]Car, error) {
	var err error
	cars := []Car{}
	err = db.Debug().Model(&Car{}).Find(&cars).Error
	if err != nil {
		return &[]Car{}, err
	}
	return &cars, nil
}

func (c *Car) FindCarByID(db *gorm.DB, id string) (*Car, error) {
	var err error
	err = db.Debug().Model(Car{}).Where("id = ?", id).Take(&c).Error
	if err != nil {
		return &Car{}, err
	}
	return c, nil
}

func (c *Car) UpdateCar(db *gorm.DB) (*Car, error) {
	err := db.Debug().Model(&Car{}).Where("id = ?", c.ID).Take(&Car{}).Updates(
		Car{
			Model:         c.Model,
			Brand:         c.Brand,
			Year:          c.Year,
			NumberOfDoors: c.NumberOfDoors,
			NumberOfSeats: c.NumberOfSeats,
			UpdatedAt:     time.Now(),
		},
	).Error
	if err != nil {
		return &Car{}, err
	}
	return c, nil
}

func (c *Car) DeleteCar(db *gorm.DB, id string) error {
	err := db.Debug().Model(&Car{}).Where("id = ?", id).Take(&Car{}).Delete(&Car{}).Error
	if err != nil {
		return err
	}
	return nil
}
