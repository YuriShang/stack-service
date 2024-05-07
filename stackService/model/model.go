package model

import (
	"gorm.io/gorm"
)

type StackData struct {
	gorm.Model
	Id   uint `gorm:"not null;unique;autoIncrement:true"`
	Data int  `gorm:"type:int;not null"`
}
