package interfaces

import "gorm.io/gorm"

type Database interface {
	GetInstance() *gorm.DB
}
