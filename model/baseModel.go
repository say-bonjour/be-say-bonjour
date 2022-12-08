package model

import (
	"time"
)

type BaseModel struct {
	Id uint64 `gorm:"primary_key" json:"id"`

	CreatorId uint64    `gorm:"type:bigint(20) unsigned" json:"creatorId"`
	CreatedAt time.Time `json:"createdAt"`

	UpdaterId uint64    `gorm:"type:bigint(20) unsigned" json:"updaterId"`
	UpdatedAt time.Time `json:"updatedAt"`
}
