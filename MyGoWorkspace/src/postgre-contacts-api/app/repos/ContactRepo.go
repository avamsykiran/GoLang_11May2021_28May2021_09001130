package repos

import (
	"database/sql"
	"fmt"

	"postgre-contacts-api/app/models"

	_ "github.com/lib/pq"
)

func AddContact(c models.Contact) models.Contact {
	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	insStmt := "insert into contacts values($1,$2,$3,$4,$5,$6)"

	res, err := conn.Exec(insStmt, c.ContactId, c.FirstName, c.LastName, c.Mobile, c.AlternateMobile, c.MailId)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d rows added successfully", count)
	return c
}

func UpdateContact(c models.Contact) models.Contact {
	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	updStmt := "update dept set first_name=$1 ,last_name=$2,mobile=$3 ,alternate_mobile=$4 ,mail_id=$5 where contact_id=$6"

	res, err := conn.Exec(updStmt, c.FirstName, c.LastName, c.Mobile, c.AlternateMobile, c.MailId, c.ContactId)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d rows updated successfully", count)

	return c
}

func DelContact(cid int) {
	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	delStmt := "delete from contacts where contact_id=$1"

	res, err := conn.Exec(delStmt, cid)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d rows deleted successfully", count)
}

func GetContactById(cid int) models.Contact {

	var c models.Contact

	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	selStmt := "select contact_id,first_name,last_name,mobile ,alternate_mobile,mail_id from contacts where contact_id=$1"

	row := conn.QueryRow(selStmt, cid)

	err = row.Scan(&c.ContactId, &c.FirstName, &c.LastName, &c.Mobile, &c.AlternateMobile, &c.MailId)

	if err != nil {
		panic(err)
	}

	return c
}

func GetAllContacts() []models.Contact {

	var contacts []models.Contact

	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	selStmt := "select contact_id,first_name,last_name,mobile ,alternate_mobile,mail_id from contacts"

	rows, err := conn.Query(selStmt)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var c models.Contact
		row.Scan(&c.ContactId, &c.FirstName, &c.LastName, &c.Mobile, &c.AlternateMobile, &c.MailId)
		contacts = append(contacts, c)
	}

	return contacts
}
