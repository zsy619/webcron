package models

import (
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost, _ := beego.AppConfig.String("db.host")
	dbport, _ := beego.AppConfig.String("db.port")
	dbuser, _ := beego.AppConfig.String("db.user")
	dbpassword, _ := beego.AppConfig.String("db.password")
	dbname, _ := beego.AppConfig.String("db.name")
	timezone, _ := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(User), new(Task), new(TaskGroup), new(TaskLog))

	runMode, _ := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	tp, _ := beego.AppConfig.String("db.prefix")
	return tp + name
}
