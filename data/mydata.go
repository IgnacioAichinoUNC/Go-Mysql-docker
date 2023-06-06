package data

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	//"log"
	"fmt"
	"Aichino/dockergo/model"
)


var database *sql.DB

func ConnectionString() error {
	
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
    if err != nil {
        return err
    }
	database = db

	return nil
}


func GetAllUsers() [] model.User {

	selectQuery := "SELECT * FROM usuarios;"
	rows, err := database.Query(selectQuery)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}


func Insertnewuser( adduser model.User) error {

	safepass, err := bcrypt.GenerateFromPassword([]byte(adduser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO usuarios (username, password)
					VALUES(? , ?)`

	stmt, err := database.Prepare(insertQuery)
	if err != nil {
        return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(adduser.Username, safepass)
	if err != nil {
        fmt.Println("Error al ejecutar la consulta:", err)
		return err
	}
	fmt.Println("Inserción exitosa database")
	return nil
}

func GetUser( user model.User) model.User {

	selectQuery := "SELECT * FROM usuarios WHERE USERNAME=?;"
	row := database.QueryRow(selectQuery, user.Username)

	var user_result  model.User
	var id int
	err := row.Scan(&id, &user_result.Username , &user_result.Password)

	if err != nil {
		println("Error select username login")

	}

	return user_result
}

