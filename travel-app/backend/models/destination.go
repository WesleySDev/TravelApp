// Tabela banco de dados
package models

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Name        string `gorm:"size:100;not null"`
	Country     string `gorm:"size:100;not null"`
	Description string `gorm:"type:text;not null"`
	ImageURL    string `gorm:"size:255"`

	Packages []Package `gorm:"foreignKey:DestinationID"`
}
