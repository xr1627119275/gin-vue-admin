package global

import (
	"github.com/gofrs/uuid/v5"
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID ,.
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type GVA_MODEL_UUID struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primary_key"` // 主键ID ,.
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func (u *GVA_MODEL_UUID) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, _ = uuid.NewV4()
	return
}
