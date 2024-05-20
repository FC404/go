package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// 26.177.71.172:80
func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/user/insert", IndexUser)
	http.HandleFunc("/user/delUser", delUser)
	http.HandleFunc("/user/upUser", upUser)
	http.HandleFunc("/user/ckUser", checkUser)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	msg := "欢迎光临我的首页"
	w.Write([]byte(msg))
}

// 添加
func IndexUser(w http.ResponseWriter, r *http.Request) {
	//接受客户端提交的数据
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

	//数据库连接
	ConnString := "root:123456@tcp(127.0.0.1:3306)/mzy?charset=utf8mb4"
	db, err := sql.Open("mysql", ConnString)
	if err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭
	defer db.Close()

	//定义SQL语句
	sqlString := "insert into users (username,password) values (?,?)"

	//创建预编译对象
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	//延迟关闭 stmt
	defer stmt.Close()

	result, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	lastId, _ := result.LastInsertId()
	//定义要输出字符串信息
	msg := fmt.Sprintf("成功添加一条记录，记录编号为%v", lastId)
	title := "添加咯\n"
	w.Write([]byte(title))
	//向客户端输出指定信息
	w.Write([]byte(msg))
}

// 删除
func delUser(w http.ResponseWriter, r *http.Request) {
	//接受客户端提交的数据
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

	//数据库连接
	ConnString := "root:123456@tcp(127.0.0.1:3306)/mzy?charset=utf8mb4"
	db, err := sql.Open("mysql", ConnString)
	if err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭
	defer db.Close()

	//定义SQL语句
	sqlString := "DELETE FROM users WHERE username = ? AND password = ?"

	//创建预编译对象
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	//延迟关闭 stmt
	defer stmt.Close()

	result, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	lastId, _ := result.RowsAffected()
	//定义要输出字符串信息
	msg := fmt.Sprintf("成功删除，记录编号为%v", lastId)
	title := "删除咯\n"
	w.Write([]byte(title))
	//向客户端输出指定信息
	w.Write([]byte(msg))
}

// 修改
func upUser(w http.ResponseWriter, r *http.Request) {
	//接受客户端提交的数据
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

	//数据库连接
	ConnString := "root:123456@tcp(127.0.0.1:3306)/mzy?charset=utf8mb4"
	db, err := sql.Open("mysql", ConnString)
	if err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭
	defer db.Close()

	//定义SQL语句
	sqlString := "UPDATE users SET password=? WHERE username=?"

	//创建预编译对象
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	//延迟关闭 stmt
	defer stmt.Close()

	result, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	lastId, _ := result.RowsAffected()
	//定义要输出字符串信息
	msg := fmt.Sprintf("成功修改，记录编号为%v", lastId)
	title := "修改咯\n"
	w.Write([]byte(title))
	//向客户端输出指定信息
	w.Write([]byte(msg))
}

// 查询
func checkUser(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Id       int
		Username string
		Password string
	}
	//接受客户端提交的数据
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

	//数据库连接
	ConnString := "root:123456@tcp(127.0.0.1:3306)/mzy?charset=utf8mb4"
	db, err := sql.Open("mysql", ConnString)
	if err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭
	defer db.Close()

	//定义SQL语句
	sqlString := "UPDATE users SET password=? WHERE username=?"

	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println(err)
		return
	}

	//把数据放进切片
	var users []User
	for rows.Next() {
		//结构体user
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(user)
		temp := append(users, user)
		users := temp
		fmt.Println(users)
	}

	//创建预编译对象
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	//延迟关闭 stmt
	defer stmt.Close()

	result, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println(666, err)
		return
	}
	lastId, _ := result.RowsAffected()

	//定义要输出字符串信息
	msg := fmt.Sprintf("成功修改，记录编号为%v", lastId)
	title := "查询咯\n"
	w.Write([]byte(title))
	//向客户端输出指定信息
	w.Write([]byte(msg))
}
