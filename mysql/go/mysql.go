package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)



//数据库配置
const (
	userName = "root"
	password = ""
	ip = "127.0.0.1"
	port = "3306"
	dbName = "demo"
)



type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}




//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func CloseDB(){
	defer DB.Close()
}


func getAll()([]User){
	rows,err:=DB.Query("select * from user")
	if err!=nil{
		fmt.Println("查询出错了")
	}

	var users []User
	for rows.Next(){
		var user User
		err:=rows.Scan(&user.Id,&user.Username,&user.Password)
		if err!=nil{
			fmt.Println("rows fail")
		}

		//将user追加到users的这个数组中

		users = append(users,user)
	}

	return users
}


func del(arrId []string)(id int64, err error){
	idStr := strings.Join(arrId, "','")
	sql := "DELETE FROM user WHERE id in ('%s')"

	sqlText := fmt.Sprintf(sql, idStr)
	result, err := DB.Exec(sqlText)
	if err != nil{
		fmt.Println("del failed")
		return 0,err
	}

	fmt.Println("delete data successd:", result)

	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return 0,err
	}

	fmt.Println("Affected rows:", rowsaffected)
	return rowsaffected,err
}


func insert(user *User)(id int64, err error){
	//准备sql语句
	stmt, err := DB.Prepare("INSERT INTO user (`username`, `password`) VALUES (?, ?)")
	if err != nil{
		fmt.Println("Prepare fail")
		return 0,err
	}
	//将参数传递到sql语句中并且执行
	result, err := stmt.Exec(user.Username, user.Password)
	if err != nil{
		fmt.Println("Exec fail")
		return 0,err
	}


	lastInsertID,err := result.LastInsertId()    //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return 0,err
	}
	fmt.Println("Insert data id:", lastInsertID)

	rowsaffected,err := result.RowsAffected()  //通过RowsAffected获取受影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return 0,err
	}
	fmt.Println("Affected rows:", rowsaffected)

	return  lastInsertID,err ;
}


func update(user *User) (rows int64, err error){
	//准备sql语句
	stmt, err := DB.Prepare("UPDATE user SET username = ?, password = ? WHERE id = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return 0,err
	}
	//设置参数以及执行sql语句
	result, err := stmt.Exec(user.Username, user.Password, user.Id)
	if err != nil{
		fmt.Println("Exec fail")
		return 0,err
	}

	fmt.Println("update data successd:", result)

	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return 0,err
	}
	fmt.Println("Affected rows:", rowsaffected)
	return rowsaffected,err
}


func main(){


	InitDB()

	list:= User{
		Username:"liu",
		Password:"das",
	}
	insertId,_:=insert(&list)

	fmt.Println(insertId)


	set:= User{
		Username:"ak",
		Password:"ak",
		Id:12,
	}
	rowsaffected,_:= update(&set)

	fmt.Println(rowsaffected)



	res:= getAll()
	for _,item :=range res{
		fmt.Printf("%+v\n", item)
	}



	//arrId := []string{"1","2","3"}
	//result := del(arrId)
	//fmt.Println("%+v\n", result)

	CloseDB()


}


