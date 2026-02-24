// Tabela banco de dados
package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	Destination uint `gorm:"not null"`
}
