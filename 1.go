package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbConn struct {
	Dsn string
	Db  *sql.DB
}

func main() {
	var err error
	dbConn := DbConn{
		Dsn: "root:123456@tcp(127.0.0.1:3306)/mzy?charset=utf8mb4",
	}
	dbConn.Db, err = sql.Open("mysql", dbConn.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbConn.Db.Close()
	preExecData(&dbConn)

	// connString := "root:123456@tcp(127.0.0.1:3306)/mzy?charset=utf8mb4"
	// db, err := sql.Open("mysql", connString)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// preExecData(db)
	// defer db.Close()
	//添加
	// result, err := db.Exec("insert into users (username,password) values (?,?)", "jerry", "123")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(id)

	//修改
	// result, err := db.Exec("update users set password=? where id=?", "404", 3)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// num, err := result.RowsAffected()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(num)

	//删除
	// result, err := db.Exec("delete  from  users where password=?", "404")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// num, err := result.RowsAffected()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(num)

	//查询
	// type User struct {
	// 	Id       int
	// 	Username string
	// 	password string
	// }

	// rows, err := db.Query("select * from users")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //结构体user
	// var user User
	// //把数据放进切片
	// var users []User
	// for rows.Next() {
	// 	err = rows.Scan(&user.Id, &user.Username, &user.password)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// fmt.Println(user)
	// 	users = append(users, user)
	// }
	// fmt.Println(users)
}

func preExecData(dbConn *DbConn) {
	//添加
	result, err := dbConn.Db.Exec("insert into users (username,password) values (?,?)", "jerry", "123")
	//删除
	//result, err := dbConn.Db.Exec("delete  from  users where password=?", "404")
	//查找
	// rows, err := dbConn.Exec("select * from users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Last Insert ID:", lastInsertId)
	fmt.Println("Rows Affected:", rowsAffected)
}
