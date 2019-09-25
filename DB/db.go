package db

import (
	"context"
	"fmt"
	"log"

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

//var DBB *sql.DB
var InsertData *sql.Stmt

func DB_connectMySQL() *sql.DB {
	// Create the database handle, confirm driver is present
	fmt.Println("Connect Database...")
	DBB, err := sql.Open("mysql", "trangx:@tcp(127.0.0.1:3306)/shortenurl")

	if err != nil {
		fmt.Println(err)
	}
	//insert, err := DBB.Query("INSERT INTO test VALUES ( 'fulllll', 'shortttt' )")
	//fmt.Println(insert)
	// InsertData, es := DBB.Prepare("insert into test values(?,?)")
	// if es != nil {
	// 	log.Println(es, InsertData)
	// }
	return DBB
}

func GetCollection(name string) *mongo.Collection {
	Col := DB.Collection(name)
	return Col
}
