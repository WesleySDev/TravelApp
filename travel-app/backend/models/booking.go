package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PackageID uint `gorm:"not null"`
	Package Package `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Status string `gorm:"type:text;default:'pending';check:status IN ('pending', 'confirmed', 'cancelled')"`
}  