package links

import (
	"context"
	"log"
	"strconv"

	database "github.com/TeddyMuli/go_graphql_api/internal/pkg/db/psql"
	"github.com/TeddyMuli/go_graphql_api/internal/users"
)

// #1
type Link struct {
	ID      string
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
	rows, err := database.Db.Query(context.Background(), "SELECT id, title, address FROM Links")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var links []Link
	var username string
	var id int

	for rows.Next() {
		var link Link
		if err := rows.Scan(&id, &link.Title, &link.Address, &id, &username); err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       strconv.Itoa(id),
			Username: username,
		}
		links = append(links, link)
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}

	return links
}
