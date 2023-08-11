package main

import (
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type User struct {
	Id   string
	Name string
}

func main() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db connection")
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping db")
	}

	db.SetMaxOpenConns(10)

	//q := `INSERT INTO users (id, username) VALUES ('d057a64f-9819-4fb1-b09a-0b98bb72ec02', 'testUser3');`
	//
	//rows, err := db.Query(q)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("failed to query db")
	//}

	rows, err := db.Query("SELECT username FROM users WHERE id = $1", "d057a64f-9819-4fb1-b09a-0b98bb72ec00")

	if err != nil {
		log.Fatal().Err(err).Msg("failed to query db")
	}

	var usernames []string

	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			log.Fatal().Err(err).Msg("failed to scan row")
		}

		usernames = append(usernames, username)
		log.Info().Str("username", username).Msg("found user")
	}

	fmt.Println(len(usernames), usernames)

	var u User

	s := sqlbuilder.Select("id", "username").From("users").Where("id = 'd057a64f-9819-4fb1-b09a-0b98bb72ec02'").String()

	fmt.Println(s)

	err = db.QueryRow(s).Scan(&u.Id, &u.Name)
	if err != nil {
		log.Fatal().Err(err).Msg("failed")
	}

	fmt.Println(u)
}
