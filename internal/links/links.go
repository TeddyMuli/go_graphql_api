package links

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	database "github.com/TeddyMuli/go_graphql_api/internal/pkg/db/psql"
	"github.com/TeddyMuli/go_graphql_api/internal/users"
)

// #1
type Link struct {
	ID      int
	Title   string
	Address string
	User    *users.User
}

//#2
func (link Link) Save() int64 {
	query := "INSERT INTO Links(Title, Address, UserID) VALUES($1, $2, $3) RETURNING ID"

	var id int64
	err := database.Db.QueryRow(context.Background(), query, link.Title, link.Address, link.User.ID).Scan(&id)
	if err != nil {
		log.Printf("InsertLink error: %v", err)
		return 0
	}

	log.Println("Row inserted!")
	return id
}

func GetAll() []Link {
	rows, err := database.Db.Query(context.Background(), "SELECT l.ID, l.Title, l.Address, u.ID, u.Username FROM Links l LEFT JOIN Users u ON l.UserID = u.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var links []Link

	for rows.Next() {
		var link Link
		var username sql.NullString
		var UserID int

		if err := rows.Scan(&link.ID, &link.Title, &link.Address, &UserID, &username); err != nil {
			log.Printf("Scan error: %v", err)
      continue
		}

		if username.Valid {
			link.User = &users.User{
				ID:       UserID,
				Username: username.String,
			}
		}
		links = append(links, link)
	}

	if rows.Err() != nil {
		log.Printf("Rows iteration error: %v", rows.Err())
    return []Link{}
	}

	return links
}
