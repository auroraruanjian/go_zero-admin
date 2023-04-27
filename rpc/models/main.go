package main

import (
	"go-zero-demo/rpc/common/helper"
	"go-zero-demo/rpc/models/adminmodel"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Dynamic SQL
type UserQuerier interface {
	// SELECT * FROM @@table WHERE name = @name
	FindUserByName(name string) (*gen.T, error)
}

type PermissionQuerier interface {
	//SELECT distinct * FROM @@table WHERE id IN(
	//	SELECT id FROM admin_role_permission WHERE admin_role_permission.admin_role_id IN(@role)
	//)
	FindByRoleId(role string) ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:root@(192.168.1.61:3306)/go-zero?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(adminmodel.AdminUser{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User`
	g.ApplyInterface(func(UserQuerier) {},
		adminmodel.AdminUser{},
	)

	g.ApplyInterface(func() {},
		adminmodel.AdminRole{},
	)

	g.ApplyInterface(func(PermissionQuerier) {},
		adminmodel.AdminPermission{},
	)

	// Generate the code
	g.Execute()

	gormdb.Migrator().DropTable(
		adminmodel.AdminUser{},
		adminmodel.AdminRole{},
		adminmodel.AdminPermission{},
	)
	gormdb.AutoMigrate(
		adminmodel.AdminUser{},
		adminmodel.AdminRole{},
		adminmodel.AdminPermission{},
	)

	// 写入预置数据
	password, _ := helper.EncryptPassword("123456")
	user := &adminmodel.AdminUser{
		Name:     "admin",
		NickName: "administrator",
		Password: password,
		Status:   1,
		CreateBy: "system",
		AdminRole: []*adminmodel.AdminRole{
			{
				Name: "admin",
				Slug: "administrator",
			},
		},
	}
	gormdb.Clauses(clause.OnConflict{
		//Columns:   []clause.Column{{Name: "id"}},
		//DoUpdates: clause.AssignmentColumns([]string{"name", "nick_name", "password"}),
		UpdateAll: true,
	}).Create(user)
}
