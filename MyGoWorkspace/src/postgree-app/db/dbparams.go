package db

import "fmt"

const (
	Host     = "localhost"
	Port     = 5432
	Username = "postgres"
	Password = "root"
	Db       = "hrdb"
)

func ConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, Username, Password, Db)
}
