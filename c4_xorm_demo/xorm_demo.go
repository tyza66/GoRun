package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

/**
 * Author: tyza66
 * Date: 2023/04/20 8:19
 * Github: https://github.com/tyza66
 **/
func main() {
	var (
		userName  string = "root"
		password  string = "123456"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	//这个结构体用于和数据库中的表同步 修改这个表也会变
	type User struct {
		Id      int64
		Name    string
		Salt    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}
	//同步数据库中的表
	err = engine.Sync(new(User))
	if err != nil {
		fmt.Println("同步表结构失败")
	}

	//插入使用insert 可以同时插入多个数据
	user := User{Id: 10001, Name: "giao", Age: 18}
	//第一个返回值是影响的行数 这里我们一般使用的是指针的参数
	insert, err := engine.Insert(&user)
	if err != nil {
		fmt.Println("插入失败")
	}
	fmt.Println(insert)
	//一次提交多个
	user1 := User{Id: 10003, Name: "sdm", Age: 18}
	user2 := User{Id: 10004, Name: "ll", Age: 18}
	insert,_ = engine.Insert(user1, user2)
	if insert>=1 {
		fmt.Println("插入成功")
	}

	//使用切片提交
	var users []User
	users = append(users, User{Id: 10005, Name: "sdm2", Age: 18})
	users = append(users, User{Id: 10006, Name: "sdm3", Age: 18})
	n,_ = engine.Insert(users)
	if n>=1 {
        fmt.Println("插入成功")
    }
}
