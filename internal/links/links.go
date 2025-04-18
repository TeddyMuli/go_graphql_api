package links

import (
	"context"
	"log"

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
	query := "INSERT INTO Links(Title, Address) VALUES($1, $2) RETURNING ID"

	var id int64
	err := database.Db.QueryRow(context.Background(), query, link.Title, link.Address).Scan(&id)
	if err != nil {
		log.Printf("InsertLink error: %v", err)
		return 0
	}

	log.Println("Row inserted!")
	return id
}
