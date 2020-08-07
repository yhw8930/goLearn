package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义一个数据模型(user表)
// 列名是字段名的蛇形小写(PassWd->pass_word)
type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

func main() {
	db, err := gorm.Open("mysql", "root:yhw8930@tcp(120.79.255.48:3306)/mytest?&charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
	}
	defer db.Close()
	// 自动迁移数据结构(table schema)
	// 注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
	db.SingularTable(true) //可以取消表名的复数形式，使得表名和结构体名称一致
	/*db.AutoMigrate(&User{})
	db.AutoMigrate(&User{}, &Product{}, &Order{})
	// 创建表时添加表后缀
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})*/

	// 为模型`User`创建表
	db.CreateTable(&User{})

	// 检查模型`User`表是否存在
	fmt.Println(db.HasTable(&User{}))
	// 检查表`users`是否存在
	fmt.Println(db.HasTable("user"))

	// 删除模型`User`的表
	//db.DropTable(&User{})
	// 删除表`users`
	db.DropTable("user")
}
