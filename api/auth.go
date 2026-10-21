package api

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type USERLOGIN struct {
	Username string `json:"username" binding="required"`
	Password string `json:"password" binding="required"`
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
func Login(c *gin.Context) {
	var userlogin USERLOGIN
	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		panic(err)
	}
	if err := c.ShouldBindJSON(&userlogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

	println(userlogin.Username)
	println(userlogin.Password)
	if userlogin.Username == "technik" && userlogin.Password == "tt@mbg" {
		c.Status(http.StatusAccepted)
		cookie, _ := randomHex(20)
		sqlStatement := `INSERT INTO sessions_users (sessionid) VALUES ($1)`
		results, err := db.Exec(sqlStatement, cookie)
		println(results)
		if err != nil {
			panic(err)

		}
		c.Status(http.StatusAccepted)
		c.SetCookie("sessionid", cookie, 80000, "/", "localhost", false, true)
	} else {
		c.Status(http.StatusForbidden)
	}

}

type USERCHECK struct {
	Sessionid string `json:"sessionid" binding="required"`
}

func CheckSession(c *gin.Context) {

	var usercheck USERCHECK

	PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PsqlInfo)
	sqlStatement := `SELECT sessionid FROM WHERE song_id=$1`
	if err != nil {
		panic(err)
	}
	if err := c.ShouldBindJSON(&usercheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

}
