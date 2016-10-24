//前端采用的是开源的AdminLTE-2.3.6
//系统的部分代码直接使用了beego作者的beeweb的代码，在此表示感谢
package main

import (
	_ "pms/db"
	. "pms/init"
	// . "pms/models/base"
	_ "pms/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	APP_VER = "1.0.0"
)

//初始化函数，会在main之前执行，可用于初始化数据，连接数据库等操作
func init() {
	dbType := beego.AppConfig.String("db_type")
	//获得数据库参数，不同数据库可能存在没有值的情况没有的值nil
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	orm.RegisterDriver(dbType, orm.DRPostgres)

	switch dbType {
	//数据库类型和数据库驱动名一致
	case "postgres":
		dbSslmode := beego.AppConfig.String(dbType + "::db_sslmode")
		dataSource := "user=" + dbUser + " password=" + dbPwd + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=" + dbSslmode
		orm.RegisterDataBase(dbAlias, dbType, dataSource)

	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "db_charset")
		dataSource := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=" + dbCharset
		orm.RegisterDataBase(dbAlias, dbType, dataSource)

	}
	//重新运行时是否覆盖原表创建,false:不会删除原表,修改表信息时将会在原来的基础上修改，true删除原表重新创建
	coverDb, _ := beego.AppConfig.Bool("cover_db")

	//自动建表
	orm.RunSyncdb(dbAlias, coverDb, true)
	InitApp()
	InitDb()
	// filter.FilterInit()
	beego.AddFuncMap("i18n", i18n.Tr)

}
func main() {
	// var user User
	// user.Name = "admin"
	// user.Mobile = "123123"
	// o := orm.NewOrm()
	// o.Using("default")
	// id, err := o.Insert(&user)
	// fmt.Println(id)
	// fmt.Println(err)
	beego.Run()
}
