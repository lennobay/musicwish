package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Songentry struct {
	Id     int
	Songid string
	Votes  int
}

func RemoveSong(c *gin.Context) {
	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM eventtable_" + c.Query("event") + " WHERE id=" + c.Query("id"))
	if err != nil {
		panic(err)
	}
	c.Status(http.StatusOK)
}

func AllSongsOfEvent(c *gin.Context) {
	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT * FROM eventtable_` + c.Query("event")

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var songs []Songentry

	for rows.Next() {
		var song Songentry
		if err := rows.Scan(&song.Id, &song.Songid, &song.Votes); err != nil {
			panic(err)
		}
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, songs)
}

func AddNewSong(c *gin.Context) {

	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}
	songid := c.DefaultQuery("song_id", "none-available")
	event_name := c.DefaultQuery("event_name", "none-available")
	if songid == "none-available" || event_name == "none-available" {
		c.String(http.StatusBadRequest, "No Song ID given or Event given")
	} else if songid == "none-available" && event_name == "none-available" {
		c.String(http.StatusBadRequest, "No Song ID given or Event given")
	} else {
		sqlStatement := `INSERT INTO eventtable_` + c.Query("event_name") + `(song_id, song_votes) values($1, 0)`
		_, err = db.Exec(sqlStatement, c.Query("song_id"))
		if err != nil {
			c.Status(http.StatusAccepted)

		} else {
			sqlStatement := `SELECT id, song_votes FROM eventtable_` + c.Query("event_name") + ` WHERE song_id=$1`
			row := db.QueryRow(sqlStatement, c.Query("song_id"))
			var id int
			var song_votes int
			switch err := row.Scan(&id, &song_votes); err {
			case sql.ErrNoRows:
				fmt.Println("No rows were returned")
			case nil:
				fmt.Println(song_votes)
			default:
				panic(err)
			}

			sqlStatementUpdateVote := `UPDATE eventtable_` + c.Query("event_name") + ` SET song_votes=$1 WHERE id=$2`
			_, err = db.Exec(sqlStatementUpdateVote, song_votes+1, id)
			if err != nil {
				panic(err)
			}
		}

	}
	defer db.Close()
}
func VoteforSong(c *gin.Context) {
	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT id, song_votes FROM eventtable_` + c.Query("event_name") + ` WHERE song_id=$1`
	row := db.QueryRow(sqlStatement, c.Query("song_id"))
	var id int
	var song_votes int
	switch err := row.Scan(&id, &song_votes); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Println(song_votes)
	default:
		panic(err)
	}

	sqlStatementUpdateVote := `UPDATE eventtable_` + c.Query("event_name") + ` SET song_votes=$1 WHERE id=$2`
	_, err = db.Exec(sqlStatementUpdateVote, song_votes+1, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
