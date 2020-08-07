package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday time.Time
	//Email       string      `gorm:"type:varchar(50)"`
	//PassWord    string      `gorm:"type:varchar(25)"`
}

func main() {
	db := getDatabase()
	defer db.Close()
	user := User{Name: "Jin66zhu", Age: 18, Birthday: time.Now()}
	//db.CreateTable(&User{})
	fmt.Println(db.NewRecord(user)) // => 主键为空返回`true`
	//db.Create(&user)
	fmt.Println(db.NewRecord(user)) // => 创建`user`后返回`false`
	// 获取第一条记录，按主键排序
	fmt.Println(db.First(&User{}).Value)
	// SELECT * FROM user ORDER BY id LIMIT 1;
	fmt.Println(db.Last(&User{}).Value)
	// SELECT * FROM user ORDER BY id DESC LIMIT 1;
	// 获取所有记录
	var users []User
	//fmt.Println(db.Find(&users).Value)
	// SELECT * FROM user;
	// 使用主键获取记录
	fmt.Println(db.First(&User{}, 10).Value)
	// 获取第一个匹配记录
	//fmt.Println(db.Where("name = ?", "jinzhu").First(&User{}).Value)
	// SELECT * FROM user WHERE name = 'jinzhu' limit 1;
	// 获取所有匹配记录
	//fmt.Println(db.Where("name = ?", "jinzhu").Find(&users).Value)
	// SELECT * FROM users WHERE name = 'jinzhu';
	fmt.Println(db.Where("name <> ?", "jinzhu").Find(&users).Value)
	db.Where("name <> ?", "jinzhu").Find(&users)

	/*// IN
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// Time
	db.Where("updated_at > ?", lastWeek).Find(&users)
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)*/
	//Where查询条件 (Struct & Map)
	//注意：当使用struct查询时，GORM将只查询那些具有值的字段
	/*// Struct
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键的Slice
	db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);*/

	//Not条件查询
	/*db.Not("name", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;

	// Not In
	db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Not In slice of primary keys
	db.Not([]int64{1,2,3}).First(&user)
	//// SELECT * FROM users WHERE id NOT IN (1,2,3);

	db.Not([]int64{}).First(&user)
	//// SELECT * FROM users;

	// Plain SQL
	db.Not("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE NOT(name = "jinzhu");

	// Struct
	db.Not(User{Name: "jinzhu"}).First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu";*/

}

func getDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:yhw8930@tcp(120.79.255.48:3306)/mytest?&charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
	}
	db.SingularTable(true)
	return db
}
