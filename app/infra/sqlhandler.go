package infra

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type SqlHandler struct {
	Conn *sql.DB
}


func NewSqlHandler() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	user := os.Getenv("MYSQL_USER")
	pass:= os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	var path string = fmt.Sprintf("%s:%s@tcp(db)/%s?charset=utf8&parseTime=true",user, pass, database)

	conn, err := sql.Open("mysql", path)

	if err != nil {
		fmt.Println(err)
	}

	return conn

}