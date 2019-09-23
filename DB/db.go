package db

import (
	"context"
	"fmt"
	"log"

	//"reflect"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func DB_connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("shortenurl")
	fmt.Println("Connected to MongoDB!")
}

var DBB *sql.DB

func DB_connectMySQL() {
	// Create the database handle, confirm driver is present
	DBB, err := sql.Open("mysql", "trang:@tcp(127.0.0.1:3306)/shortenurl")

	if err != nil {
		fmt.Println(err)
	}
	insert, err := DBB.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	fmt.Println(insert)
	defer DBB.Close()

	// Connect and check the server version
	var version string
	DBB.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

func GetCollection(name string) *mongo.Collection {
	Col := DB.Collection(name)
	return Col
}
