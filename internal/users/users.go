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
	var id int
	query := "INSERT INTO Users(Username,Password) VALUES($1,$2) RETURNING ID"
	err = database.Db.QueryRow(context.Background(), query, user.Username, hashedPassword).Scan(&id)
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

func GetUserIdByUsername(username string) (int, error) {
	var Id int

	err := database.Db.QueryRow(context.Background(), "select ID from Users WHERE Username = $1", username).Scan(&Id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return Id, nil
}

func (user *User) Authenticate() bool {
	var hashedPassword string
	err := database.Db.QueryRow(context.Background(), "select Password from Users WHERE Username = $1", user.Username).Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}
