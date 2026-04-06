package models
import "gorm.io/gorm"

type Favorite struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DestinationID uint		`gorm:"not null"`
	Destination Destination `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`


}