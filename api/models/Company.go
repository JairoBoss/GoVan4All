package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Company struct {
	ID          string    `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Address     string    `gorm:"not null" json:"address"`
	PhoneNumber string    `gorm:"not null" json:"phone_number"`
	Email       string    `gorm:"not null" json:"email"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Company) SaveCompany(db *gorm.DB) (*Company, error) {
	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Company{}, err
	}
	return c, nil
}

func (c *Company) FindAllCompanies(db *gorm.DB) (*[]Company, error) {
	var err error
	companies := []Company{}
	err = db.Debug().Model(&Company{}).Find(&companies).Error
	if err != nil {
		return &[]Company{}, err
	}
	return &companies, nil
}

func (c *Company) FindCompanyByID(db *gorm.DB, id string) (*Company, error) {
	var err error
	err = db.Debug().Model(Company{}).Where("id = ?", id).Take(&c).Error
	if err != nil {
		return &Company{}, err
	}
	return c, nil
}

func (c *Company) UpdateCompany(db *gorm.DB) (*Company, error) {
	err := db.Debug().Model(&Company{}).Where("id = ?", c.ID).Take(&Company{}).Updates(
		Company{
			Name:        c.Name,
			Description: c.Description,
			Address:     c.Address,
			PhoneNumber: c.PhoneNumber,
			Email:       c.Email,
			UpdatedAt:   time.Now(),
		},
	).Error
	if err != nil {
		return &Company{}, err
	}
	return c, nil
}

func (c *Company) DeleteCompany(db *gorm.DB, id string) error {
	err := db.Debug().Model(&Company{}).Where("id = ?", id).Take(&Company{}).Delete(&Company{}).Error
	if err != nil {
		return err
	}
	return nil
}
