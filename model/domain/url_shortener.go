package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UrlShortener struct {
	ID          uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Slug        string         `gorm:"column:slug;unique" json:"slug"`
	Url         string         `gorm:"column:url" json:"url"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	LastVisited time.Time      `gorm:"column:last_visited;autoCreateTime" json:"last_visited"`
}

func (u *UrlShortener) BeforeCreate() (err error) {
	u.ID = uuid.New()
	return
}
