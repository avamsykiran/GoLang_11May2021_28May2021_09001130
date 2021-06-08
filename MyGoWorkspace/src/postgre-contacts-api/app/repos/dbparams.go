package repos

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "root"
	db       = "adbdb"
)

func ConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, db)
}
