package crud

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// get DB info func
func getDBinfo() string {
	data, err := os.Open("mysql/crud/dbInfo.json")
	utils.CheckErr(err)

	var info model.DBinfo
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &info)

	DBinfo := info.Username + ":" + info.Password + "@tcp(" + info.Hostname + info.Port + ")/" + info.Database
	return DBinfo
}

// ConnectDB func
func ConnectDB() *sql.DB {
	DBinfo := getDBinfo()

	db, err := sql.Open("mysql", DBinfo)
	utils.CheckErr(err)

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

// MultiRowQueryUserList func
func MultiRowQueryUserList(query string) []model.User {
	DB := ConnectDB()
	defer DB.Close()

	var user model.User
	var allUsers []model.User

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.UID, &user.Uemail, &user.Upw, &user.Uname, &user.Ugender, &user.Uaddr, &user.Upost, &user.Uphone, &user.UjoinDate, &user.Ubirth, &user.Uagree, &user.Ulevel, &user.UjoinPath)
		utils.CheckErr(err)

		allUsers = append(allUsers, user)
	}

	return allUsers
}

// ExecQuery func
func ExecQuery(query string) {
	DB := ConnectDB()
	defer DB.Close()

	_, err := DB.Exec(query)
	utils.CheckErr(err)
}

// UpdateUser func
func UpdateUser(query string, user model.User) {
	DB := ConnectDB()
	defer DB.Close()

	_, err := DB.Exec(query, user.Uname, user.Ugender, user.Uphone, user.Ubirth, user.Uagree, user.Ulevel, user.UID)
	utils.CheckErr(err)

}

// InsertRowUserList func
func InsertRowUserList(query string, user model.User) {
	DB := ConnectDB()
	defer DB.Close()

	_, err := DB.Exec(query, user.Uemail, user.Upw, user.Uname, user.Ugender, user.Uphone, user.Ubirth, user.Uagree)
	utils.CheckErr(err)

}

// CheckUserExistInUserList func
func CheckUserExistInUserList(query string) bool {
	DB := ConnectDB()
	defer DB.Close()

	var count int

	err := DB.QueryRow(query).Scan(&count)
	if err == sql.ErrNoRows || count == 0 {
		return false
	}

	utils.CheckErr(err)
	return true

}

// GetEmailAndPwUserList func
func GetEmailAndPwUserList(query string, user model.User) bool {
	DB := ConnectDB()
	defer DB.Close()

	inputPW := user.Upw

	row := DB.QueryRow(query, user.Uemail)
	err := row.Scan(&user.Uemail, &user.Upw)
	utils.CheckErr(err)

	if inputPW == user.Upw {
		return true
	}

	return false
}
