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

	// データを挿入する処理
	article := models.Article{
		Title:    "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "saki",
	}

	const sqlStr = `
		insert into articles(title, contents, username, nice, created_at)
		values (?, ?, ?, 0, now());
	`
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 結果を確認
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
