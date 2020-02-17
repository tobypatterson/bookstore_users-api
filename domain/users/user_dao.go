package users

import (
	"github.com/tobypatterson/bookstore_users-api/datasources/mysql/users_db"
	"github.com/tobypatterson/bookstore_users-api/utils/date_utils"
	"github.com/tobypatterson/bookstore_users-api/utils/errors"
	"github.com/tobypatterson/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
)

var (
	usersDB = make(map[int64]*User)
)

// Get will retrieve a user
func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

// Save will save a user
func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId

	return nil
}
