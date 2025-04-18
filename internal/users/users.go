package users

import (
	"context"

	database "github.com/TeddyMuli/go_graphql_api/internal/pkg/db/psql"
	"golang.org/x/crypto/bcrypt"

	"log"
)

type User struct {
	ID       string `json:"id"`
	Username     string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO Users(Username,Password) VALUES($1,$2)"
	err = database.Db.QueryRow(context.Background(), query, user.Username, hashedPassword).Scan()
	if err != nil {
		log.Fatal(err)
	}

}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
