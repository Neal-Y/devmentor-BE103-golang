package database

import "time"

type SoftDeleteModel struct {
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
	IsDeleted bool       `json:"is_deleted" gorm:"default:false"`
}
