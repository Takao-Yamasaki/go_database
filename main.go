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
	articleID := 1
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	// クエリの実行
	// クエリの返り値がrowsに格納される
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// データベースから取得したデータをarticleに読み出す
	// スライスを用意
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		// rowsの中身を格納するArticle型の変数を用意
		var article models.Article
		var createdTime sql.NullTime

		// 変数にrowsの中身を読み出す
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			// スライスにappend関数で追加する
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
