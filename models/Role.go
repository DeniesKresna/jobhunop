package Models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name      string `json:"name" validate:"required"`
	CreatorID uint   `validate:"-"`

	Creator *User `gorm:"foreignKey:CreatorID;References:ID"`
}

func (b *Role) TableName() string {
	return "roles"
}

type RoleUpdate struct {
	Name      string `validate:"required"`
	CreatorID uint
}
