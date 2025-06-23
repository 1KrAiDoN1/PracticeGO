package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type Storage struct {
	DB *pgx.Conn
}

type DBset interface {
	GetUserEmail() (string, error)
	GetUserPassword() (string, error)
}
type UserDB struct {
	Id       int
	Email    string
	Password string
}

func GetUserInfo(getter DBset) (string, string, error) {
	email, err := getter.GetUserEmail()
	if err != nil {
		return "", "", err
	}

	password, err := getter.GetUserPassword()
	if err != nil {
		return "", "", err
	}

	return email, password, nil
}
func ConnectDB() (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/registration")
	if err != nil {
		log.Println(err)
	}
	return &Storage{DB: db}, nil
}

func (u *UserDB) GetUserEmail() (string, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
	}
	defer db.DB.Close(context.Background())
	var email string
	query := "SELECT email FROM users WHERE id = $1"
	err = db.DB.QueryRow(context.Background(), query, &u.Id).Scan(&email)
	if err != nil {
		return "", err
	}
	fmt.Println(email)
	return email, nil
}

func (u *UserDB) GetUserPassword() (string, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
	}
	defer db.DB.Close(context.Background())

	query := "SELECT password FROM users WHERE id = $1"
	var password string
	err = db.DB.QueryRow(context.Background(), query, &u.Id).Scan(&password)
	u.Password = password
	if err != nil {
		return "", err
	}
	fmt.Println(password)
	return password, nil
}
