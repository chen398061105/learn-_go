package main

//mysql 关系型数据库
//  ddl 操作数据库 dml 增删改查 ddl 提权
// 引擎 innodb 支持事务 整体速度快 支持表锁行锁
//myisam 查询速度快 不支持事务 只支持表锁
//下载第三方插件 go get 包路劲，默认保存在gopath的src目录下
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init
)

func main() {
	//数据库连接信息
	dsn := "root:lp881208@tcp(127.0.0.1:3306)/godb"
	db, err := sql.Open("mysql", dsn) //不会验证用户密码
	if err != nil {
		fmt.Printf("open failed err:%v", err)
		return
	}
	err = db.Ping() //验证
	if err != nil {
		fmt.Printf("db connt err:%v", err)
		return
	}
	fmt.Println("链接成功")

}
