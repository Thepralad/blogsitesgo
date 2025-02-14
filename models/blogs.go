package models

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Blog struct{
	Id int
	Title string
	Content string
	Email string
}
func RetrieveBlogs() []Blogs, error{
	DB, err := sql.Open("mysql", "root:25802580@tcp(127.0.0.1)/snet")
	if err != nil{
		return nil, err
	}

	var blogs []Blogs

	rows, err := db.Query("SELECT * FROM blogs")
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var blog Blogs
		rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Email)
		blogs = append(blogs, blog)
	}

	return blogs, nil
}