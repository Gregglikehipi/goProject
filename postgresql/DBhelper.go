package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type post struct {
	Id     int    `json:"id"`
	Data   string `json:"data"`
	UserId int    `json:"userId"`
	Likes  int    `json:"likes"`
}

type human struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func GetPost(id int) post {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from post where id = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	product := post{}

	for rows.Next() {
		p := post{}
		err := rows.Scan(&p.Id, &p.Data, &p.UserId, &p.Likes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		product = p
	}
	return product
}

func GetPosts() []post {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from post")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []post{}

	for rows.Next() {
		p := post{}
		err := rows.Scan(&p.Id, &p.Data, &p.UserId, &p.Likes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	return products
}

func InsertPost(data string, id int) {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into post (data, userId, likes) values ($1, $2, $3)",
		data, id, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}

func UpdatePost(id int, new int) {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update post set likes = $1 where id = $2", new, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество обновленных строк
}

func DeletePost(id int) {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from post where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}

func GetHumanName(name string, pass string) human {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from human where name = $1 and pass = $2", name, pass)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	product := human{}

	for rows.Next() {
		p := human{}
		err := rows.Scan(&p.Id, &p.Name, &p.Pass)
		if err != nil {
			fmt.Println(err)
			continue
		}
		product = p
	}
	return product
}

func GetHumanId(id int) human {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from human where name = $1 and pass = $2", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	product := human{}

	for rows.Next() {
		p := human{}
		err := rows.Scan(&p.Id, &p.Name, &p.Pass)
		if err != nil {
			fmt.Println(err)
			continue
		}
		product = p
	}
	return product
}

func PostHuman(name string, pass string) {

	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into human (name, pass) values ($1, $2)",
		name, pass)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}
