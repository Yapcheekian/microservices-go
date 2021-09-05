package users

import (
	"github.com/Yapcheekian/microservices-go/users/datasources/mysql/users"
	dateutils "github.com/Yapcheekian/microservices-go/users/utils/date_utils"
	"github.com/Yapcheekian/microservices-go/users/utils/errors"
	mysqlutils "github.com/Yapcheekian/microservices-go/users/utils/mysql_utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created from users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser = "DELETE FROM users WHERE id = ?;"
)

func (user *User) Delete() *errors.RestErr {
	stmt, err := users.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return mysqlutils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.Id = userId

	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := users.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}
