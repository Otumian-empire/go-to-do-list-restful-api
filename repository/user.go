package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/otumian-empire/go-to-do-list-restful-api/config"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
)

type User struct {
	*sql.DB
}

func (user *User) CreateUser(username, password string) error {
	result, err := user.Exec(CREATE_USER_QUERY, username, password)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(CREATE_USER_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(CREATE_USER_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(CREATE_USER_ERROR)
	}

	if _, err := result.LastInsertId(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(CREATE_USER_ERROR)
	}

	return nil
}

func (user *User) ReadUserById(id config.IdType) (model.User, error) {
	row := user.QueryRow(READ_USER_BY_ID_QUERY, id)

	var _user model.User
	err := row.Scan(
		&_user.Id,
		&_user.Username,
		&_user.Password,
		&_user.CreatedAt,
		&_user.UpdatedAt)

	if err != nil {
		log.Println(err)
		return model.User{}, fmt.Errorf(NO_ROW_FOUND)
	}

	return _user, nil
}

func (user *User) ReadUserByUsername(username string) (model.User, error) {
	row := user.QueryRow(GET_USERNAME_BY_USERNAME_QUERY, username)

	var _user model.User
	err := row.Scan(
		&_user.Id,
		&_user.Username,
		&_user.Password,
		&_user.CreatedAt,
		&_user.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return model.User{}, fmt.Errorf(NO_ROW_FOUND)
	}

	return _user, nil
}

func (user *User) UpdateUserPassword(id config.IdType, password string) error {
	result, err := user.Exec(UPDATE_PASSWORD_QUERY, password, id)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_USER_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_USER_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(UPDATE_USER_ERROR)
	}

	return nil
}

func (user *User) UpdateUserUsername(id config.IdType, username string) error {
	result, err := user.Exec(UPDATE_USERNAME_QUERY, username, id)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_USER_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_USER_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(UPDATE_USER_ERROR)
	}

	return nil

}

func (user *User) DeleteUser(id config.IdType) error {
	result, err := user.Exec(DELETE_USER_QUERY, id)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(DELETE_USER_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(DELETE_USER_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(DELETE_USER_ERROR)
	}

	return nil
}
