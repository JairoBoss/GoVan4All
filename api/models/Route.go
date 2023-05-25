package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Route struct {
	ID                 string    `gorm:"primary_key" json:"id"`
	UserID             uint32    `gorm:"not null" json:"user_id"`
	RouteType          string    `gorm:"not null" json:"route_type"`
	CarID              string    `gorm:"not null" json:"car_id"`
	User               User      `gorm:"foreignkey:UserID" json:"user"`
	Car                Car       `gorm:"foreignkey:CarID" json:"car"`
	NumberOfPassengers int       `gorm:"not null" json:"number_of_passengers"`
	StartLocation      string    `gorm:"not null" json:"start_location"`
	EndLocation        string    `gorm:"not null" json:"end_location"`
	StartTime          time.Time `gorm:"not null" json:"start_time"`
	EndTime            time.Time `gorm:"not null" json:"end_time"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (r *Route) SaveRoute(db *gorm.DB) (*Route, error) {
	var err error
	err = db.Debug().Create(&r).Error
	if err != nil {
		return &Route{}, err
	}
	return r, nil
}

func (r *Route) FindAllRoutes(db *gorm.DB) (*[]Route, error) {
	var err error
	routes := []Route{}
	err = db.Debug().Model(&Route{}).Find(&routes).Error
	if err != nil {
		return &[]Route{}, err
	}
	return &routes, nil
}

func (r *Route) FindRouteByID(db *gorm.DB, id string) (*Route, error) {
	var err error
	err = db.Debug().Model(Route{}).Where("id = ?", id).Take(&r).Error
	if err != nil {
		return &Route{}, err
	}
	return r, nil
}

func (r *Route) UpdateRoute(db *gorm.DB) (*Route, error) {
	err := db.Debug().Model(&Route{}).Where("id = ?", r.ID).Take(&Route{}).Updates(
		Route{
			UserID:             r.UserID,
			RouteType:          r.RouteType,
			CarID:              r.CarID,
			NumberOfPassengers: r.NumberOfPassengers,
			StartLocation:      r.StartLocation,
			EndLocation:        r.EndLocation,
			StartTime:          r.StartTime,
			EndTime:            r.EndTime,
			UpdatedAt:          time.Now(),
		},
	).Error
	if err != nil {
		return &Route{}, err
	}
	return r, nil
}

func (r *Route) DeleteRoute(db *gorm.DB, id string) error {
	err := db.Debug().Model(&Route{}).Where("id = ?", id).Take(&Route{}).Delete(&Route{}).Error
	if err != nil {
		return err
	}
	return nil
}
