package adminmodel

import (
	"go-zero-demo/rpc/common/helper"
	"time"

	"gorm.io/gorm"
)

type (
	AdminUser struct {
		Id        int          `gorm:"primaryKey;autoIncrement;type:mediumint unsigned;comment:编号"`             // 编号
		Name      string       `gorm:"column:name;uniqueIndex;type:string;not null;size:64;comment:用户名"`        // 用户名
		NickName  string       `gorm:"column:nick_name;type:string;size:64;not null;comment:昵称"`                // 昵称
		Avatar    string       `gorm:"column:avatar;type:string;size:256;not null;default:'';comment:头像"`       // 头像
		Password  string       `gorm:"column:password;type:string;size:64;not null;comment:密码"`                 // 密码
		Email     string       `gorm:"column:email;type:string;size:64;not null;default:'';comment:邮箱"`         // 邮箱
		Mobile    string       `gorm:"column:mobile;type:string;size:20;not null;default:'';comment:手机号"`       // 手机号
		Status    int          `gorm:"column:status;type:tinyint unsigned;size:2;not null;default:1;omment:状态"` // 状态  0：禁用   1：正常
		CreateBy  int          `gorm:"column:create_by;type:mediumint unsigned;not null;comment:创建人"`           // 创建人
		AdminRole []*AdminRole `gorm:"many2many:admin_role_users;"`

		CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `gorm:"type:TIMESTAMP"`
	}
)

func (u *AdminUser) BeforeCreate(tx *gorm.DB) (err error) {
	en_pwssword, err := helper.EncryptPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = en_pwssword

	return
}
