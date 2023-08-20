package repository

import (
	"database/sql"
	"fmt"
	"log"
)

// the root of the tables
type Repository struct {
	*Todo
	*User
}

// database connection
func NewRepository(driverName, dataSourceName string) (*Repository, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(GLOBAL_ERROR)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf(GLOBAL_ERROR)
	}

	return &Repository{
			User: &User{DB: db},
			Todo: &Todo{DB: db}},
		nil
}
