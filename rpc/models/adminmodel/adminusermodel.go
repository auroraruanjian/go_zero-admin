package adminmodel

import (
	"time"
)

type (
	AdminUser struct {
		Id        int64        `gorm:"primaryKey;autoIncrement;comment:编号"`                     // 编号
		Name      string       `gorm:"column:name;uniqueIndex;type:string;size:64;comment:用户名"` // 用户名
		NickName  string       `gorm:"column:nick_name;type:string;size:128;comment:昵称"`        // 昵称
		Avatar    string       `gorm:"column:avatar;type:string;size:256;comment:头像"`           // 头像
		Password  string       `gorm:"column:password;type:string;size:128;comment:密码"`         // 密码
		Email     string       `gorm:"column:email;type:string;size:64;comment:邮箱"`             // 邮箱
		Mobile    string       `gorm:"column:mobile;type:string;size:20;comment:手机号"`           // 手机号
		Status    int64        `gorm:"column:status;type:tinyint;size:2;comment:状态"`            // 状态  0：禁用   1：正常
		CreateBy  string       `gorm:"column:create_by;type:string;size:128;comment:创建人"`       // 创建人
		AdminRole []*AdminRole `gorm:"many2many:admin_role_users;"`

		CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:TIMESTAMP"`
	}
)
