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

	// クエリの定義
	articleID := 1000
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	// クエリの実行
	// クエリの返り値がrowsに格納される
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// データベースから取得したデータをarticleに読fみ出す
	var article models.Article
	var createdTime sql.NullTime

	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)
}
