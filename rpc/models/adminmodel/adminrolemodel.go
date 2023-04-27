package adminmodel

import (
	"time"
)

type (
	AdminRole struct {
		Id              int                `gorm:"primaryKey;type:smallint unsigned;autoIncrement;comment:编号"` // 编号
		Name            string             `gorm:"column:name;uniqueIndex;type:string;size:64;comment:角色名"`    // 角色名
		Slug            string             `gorm:"column:slug;uniqueIndex;type:string;size:64"`                //
		AdminUser       []*AdminUser       `gorm:"many2many:admin_role_users;"`
		AdminPermission []*AdminPermission `gorm:"many2many:admin_role_permission"`

		CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:TIMESTAMP"`
	}
)
