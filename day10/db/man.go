package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init
)

var db *sql.DB //连接池对象
var u user

func initDB() (err error) {
	//数据库连接信息
	dsn := "root:lp881208@tcp(127.0.0.1:3306)/godb"
	db, err = sql.Open("mysql", dsn) //不会验证用户密码
	if err != nil {
		return
	}
	err = db.Ping() //验证
	if err != nil {
		return
	}
	return
}

type user struct {
	id   int
	name string
	age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db err,%v", err)
		return
	}
	fmt.Println("数据库连接成功！")
	// query(1)
	// queryMore(0)
	// insert()
	// update(111, 1)
	pdo1()

}
func pdo1() {
	sql := `insert into user(name,age)values(?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"新人":  30,
		"新人2": 30,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}

}

//单行查询
func query(id int) {
	//查询单条记录的sql记录
	sql := `select id,name,age from user where id =?`
	//执行
	rowObj := db.QueryRow(sql, id)
	//拿到结果,自动关闭连接
	rowObj.Scan(&u.id, &u.name, &u.age)
	fmt.Printf("user:%#v", u)
}

//多行查询
func queryMore(id int) {
	sql := `select id,name,age from user where id > ?`
	rows, err := db.Query(sql, id)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u.id, &u.name, &u.age)
		fmt.Printf("user:%#v", u)
	}
}

//插入 更新 删除 都是exec
func insert() {
	sql := `insert into user(name,age)values("xiaoming",11)`
	ret, err := db.Exec(sql)
	if err != nil {
		return
	}
	//如果插入成功
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastId failed,err:%v", err)
		return
	}
	fmt.Println("id:", id)
}

//删除同理
func update(newAge, id int) {
	sql := `update user set age =? where id= ?`
	ret, err := db.Exec(sql, newAge, id)
	if err != nil {
		return
	}
	//如果插入成功
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("updated failed,err:%v", err)
		return
	}
	fmt.Println("更新行数：", num)
}
