package main

import (
	"database/sql"
	"fmt"
	"github.com/Takao-Yamasaki/go_database/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("connect to DB")
	// }

	const sqlStr = `
		select title, contents, username, nice from articles;
	`

	// クエリの返り値がrowsに格納される
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// スライスを用意
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		// rowsの中身を格納するArticle型の変数を用意
		var article models.Article
		// 変数にrowsの中身を読み出す
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			fmt.Println(err)
		} else {
			// スライスにappend関数で追加する
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
