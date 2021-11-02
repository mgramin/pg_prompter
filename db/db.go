package db

import (
	"fmt"
	"github.com/tg/pgpass"
	"golang.org/x/crypto/ssh/terminal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Params struct {
	Host     string
	Port     int
	User     string
	Dbname   string
	Password string
}

var CurrentParams Params

func GetPassword(Host string, Username string) string {
	password, err := pgpass.Password(Host, Username)
	if err != nil {
		panic(err)
	}
	if password == "" {
		fmt.Printf("Password for user %v:\n", Username)
		inputPassword, _ := terminal.ReadPassword(0)
		password = string(inputPassword)
	}
	return password
}

func OpenGorm() (*gorm.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		CurrentParams.Host, CurrentParams.Port, CurrentParams.User, CurrentParams.Password, CurrentParams.Dbname)
	return gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
}
