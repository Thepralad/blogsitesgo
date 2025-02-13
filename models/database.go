package models

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

// Updates the 'user' table in the database - inserts a row with values(email, password)
func UpdateDB(email string, password string) error{
	db, err := sql.Open("mysql", "root:25802580@tcp(127.0.0.1:3306)/snet")	
	if err != nil{
		return err
	}
	defer db.Close()
	
	stmt, err := db.Prepare("INSERT INTO users(email, pass) VALUES(?, ?)")
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email,password)
	if err != nil{
		return err
	}
	fmt.Println("Inserted to the database")
	return nil
}

// Retrieves a row - gets a row by using the email - returns email, password, error
func RetrieveRow(email string) (string, string, error){	
	db, err := sql.Open("mysql", "root:25802580@tcp(127.0.0.1:3306)/snet")	
	if err != nil{
		return "", "",  err
	}
	defer db.Close()
	
	stmt, err := db.Prepare("SELECT email, pass FROM users where email = ?")
	if err != nil{
		return "", "",  err
	}

	var _email, _password string
	err = stmt.QueryRow(email).Scan(&_email, &_password)
	
	fmt.Println(_email, _password)
	return _email, _password, nil 
}
