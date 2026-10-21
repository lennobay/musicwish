package api

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "musicwish"
	password = "g6k2N9pWC17X4I"
	dbname   = "musicwish"
)

func CreateNewEvent() {

	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS events   (
		id serial4 not null,
		name varchar(100) not null UNIQUE,
		eventtable varchar(100) not null,
		owner varchar(100) not null,
		constraint events_pkey primary key (id)
	)`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	sqlStatementCreateEvent := `INSERT INTO events ( name, eventtable, owner) values($1, $2, $3)`
	_, err = db.Exec(sqlStatementCreateEvent, "halloween", "eventtable_"+"halloween", "someone")
	if err != nil {
		panic(err)
	}
	event_table_name := "eventtable_" + "halloween"
	sqlStatementCreateEventTable := `	CREATE TABLE IF NOT EXISTS ` + event_table_name + ` (
		id serial4 not null,
		song_id varchar(255) not null UNIQUE,
		song_votes int not null,
		constraint` + ` eventtable_pkey_` + "halloween" + ` primary key (id)
	)`
	_, err = db.Exec(sqlStatementCreateEventTable)
	if err != nil {
		panic(err)

	}
	defer db.Close()
}

func CreateSessionDB() {

	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS sessions_users (
		id serial4 not null,
		sessionid varchar(255) not null UNIQUE
	)`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}
