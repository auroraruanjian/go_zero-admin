package main

import (
	"fmt"
	"go-zero-demo/rpc/common/helper"
	"go-zero-demo/rpc/models/adminmodel"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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

type Conf struct {
	Mysql Mysql `yaml:"Mysql"`
}

type Mysql struct {
	IP       string `yaml:"IP"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}

func main() {
	config := getConf()
	dsn := getDsn(config.Mysql)

	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(dsn))
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
	gormdb.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(
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
		CreateBy: 0,
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

func getDsn(mysqlConfig Mysql) string {
	return mysqlConfig.Username + ":" + mysqlConfig.Password + "@(" + mysqlConfig.IP + ":" + mysqlConfig.Port + ")/" + mysqlConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func getConf() Conf {
	var conf Conf // 加载文件
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	} // 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return conf
}
