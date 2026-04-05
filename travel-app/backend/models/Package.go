// Tabela banco de dados
package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model

	DestinationID uint 		`gorm:"not null"`
	Destination Destination `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Price float64 `gorm:"type:numeric(10,2);not null"`
	StartDate string `gorm:"not null"`
	EndDate string `gorm:"not null"`

	Bookings []Booking `gorm:"foreignkey:PackageID"`
}
