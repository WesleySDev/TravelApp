// Tabela banco de dados
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"size:250;not null" json: "-"`
	Role     string `gorm:"type:TEXT;default:user;check:role IN ('user','admin')"`
}
