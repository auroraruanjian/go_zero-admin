package adminmodel

import "time"

type (
	AdminPermission struct {
		Id          int64        `gorm:"primaryKey;type:smallint;autoIncrement;comment:编号"`        // 编号
		ParentId    int64        `gorm:"column:parent_id;type:smallint;comment:上级菜单ID"`            // 上级菜单ID
		Name        string       `gorm:"column:name;uniqueIndex;type:string;size:128;comment:角色名"` // 角色名
		Icon        string       `gorm:"column:icon;type:string;size:32;comment:图标"`               //图标
		Rule        string       `gorm:"column:icon;type:string;size:64;comment:权限规则"`             //权限规则
		Description string       `gorm:"column:description;type:string;size:256;comment:描述"`       //描述
		AdminRole   []*AdminRole `gorm:"many2many:admin_role_permission"`

		CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:TIMESTAMP"`
	}
)
