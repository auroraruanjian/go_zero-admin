package service

import (
	"fmt"
	"go-zero-demo/rpc/common/helper"
	"go-zero-demo/rpc/models/query"
	"go-zero-demo/rpc/sys/sysclient"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type AdminUserService struct {
	DB *gorm.DB
}

// 获取角色权限
func (a AdminUserService) GetAdminRolePermission(role []int) []*sysclient.AdminPermission {
	permission := []*sysclient.AdminPermission{}

	if len(role) == 0 {
		return permission
	}

	p := query.AdminPermission

	// 判断是否超管
	if helper.InArrayInt(1, role) {
		rows, err := p.Find()

		fmt.Println(rows, err)
	} else {
		//p.Preload(p.AdminRole).Where(p.AdminRole.,.In(1,2,3))
		//a.DB.Raw()
		role_string := []string{}
		for _, r_id := range role {
			role_string = append(role_string, strconv.Itoa(r_id))
		}
		role_list, error := p.FindByRoleId(strings.Join(role_string, ","))
		if error != nil {
			return permission
		}

		for _, role_item := range role_list {
			permission = append(permission, &sysclient.AdminPermission{
				Id:          int32(role_item.Id),
				ParentId:    int32(role_item.ParentId),
				Name:        role_item.Name,
				Icon:        role_item.Icon,
				Rule:        role_item.Rule,
				Description: role_item.Description,
			})
		}
	}

	return permission
}
