package helpers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
)

var Db *gorm.DB

func InitDB() {
	conf := GetListValue("mysql")
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//root:123456@tcp(localhost:3306)/ablog?charset=utf8
	path := strings.Join([]string{conf["username"],":",conf["password"],"@tcp(",conf["host"],":",conf["port"],")/",conf["dbname"],"?charset=utf8"},"")
	Db,_ = gorm.Open("mysql",path)
	//Db.AutoMigrate(&models.Navi{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}
