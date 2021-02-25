package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "115411"
	dbname   = "mustafayavuz"
)

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	insertStmt := "insert into users(FirstName,LastName,UserName,password,email,birthdate,isactive) values('Mustafa','Yavuz','Myavuz','123456','myavuz53@gmail.com','25-02-2021',1)"
	result, e := db.Exec(insertStmt)
	if e != nil {
		log.Fatal(e)
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted %d rows", rowCount)

	// lastID, err1 := result.LastInsertId()
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	// log.Printf("Inserted ID: %d", lastID)

	var (
		ID        int
		UserName  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		Birthdate string
		IsActive  bool
	)

	rows, err := db.Query("SELECT * FROM public.users ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&ID, &UserName, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+Birthdate+" "+strconv.FormatBool(IsActive))
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()

	data, err2 := db.Query("SELECT * FROM public.users where id=12 ORDER BY id ASC")
	if err2 != nil {
		log.Fatal(err2)
	}
	for data.Next() {
		err2 = data.Scan(&ID, &UserName, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
		if err2 != nil {
			log.Fatal(err2)
		}

		log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+Birthdate+" "+strconv.FormatBool(IsActive))
	}

	err2 = data.Err()
	if err2 != nil {
		log.Fatal(err2)
	}
	data.Close()

	// stmt, errQ := db.Prepare("SELECT * FROM public.users where id = ?")

	// if errQ != nil {
	// 	log.Fatal(errQ)
	// }

	// rows, errX := stmt.Query(1)
	// if errX != nil {
	// 	log.Fatal(errX)
	// }
	// for rows.Next() {
	// 	scanErr := rows.Scan(&ID, &UserName, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
	// 	if scanErr != nil {
	// 		log.Fatal(scanErr)
	// 	}
	// }
	// log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+UserName+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+Birthdate+" "+strconv.FormatBool(IsActive))

	// defer rows.Close()
	// defer stmt.Close()

	// stmt, err := db.Prepare("insert into users(FirstName,LastName,UserName,password,email,birthdate,isactive) values(?,?,?,?,?,?,?)")
	// UserName = "TestUSer"
	// FirstName = "Mustafa"
	// LastName = "YAvuz"
	// Email = "myavuz53@gmail.com"
	// Password = "123456"
	// Birthdate = "22-02-2021"
	// IsActive = true

	// res, errStmt := stmt.Exec(FirstName, LastName, UserName, Password, Email, Birthdate, IsActive)
	// if errStmt != nil {
	// 	log.Fatal(errStmt)
	// }
	// fmt.Println(res.LastInsertId())
	defer db.Close()
}

//CheckError hata yakalama
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
