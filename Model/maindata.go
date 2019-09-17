package maindata

import (
	"context"
	"github.com/trangnnp-ts/assignment00/db"
	"log"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Dataa struct {
	FullLink    string    `bson:"fulllink"`
	ShortenLink string    `bson:"shortenlink"`
	Used        int       `bson:"used"`
	CreatedDate time.Time `bson:"createddate"`
	UpdateDate  time.Time `bson:"updatedate"`
}

func Add(full string, short string) Dataa {
	collec := db.GetCollection("maindatas")

	if short == "" {
		short = "Trang." + RandStringBytes(9)
	}
	data := Dataa{
		FullLink:    full,
		ShortenLink: short,
		Used:        0,
		CreatedDate: time.Now(),
		UpdateDate:  time.Now()}
	insertResult, errr := collec.InsertOne(context.TODO(), &data)
	if errr != nil {
		log.Fatal(errr)
	}
	log.Fatal(insertResult)
	return data
}

func GetOne(alias string) Dataa {
	collec := db.GetCollection("maindatas")
	filter := bson.M{"$or": []bson.M{bson.M{"fulllink": alias}, bson.M{"shortenlink": alias}}}
	var p Dataa
	if err := collec.FindOne(context.TODO(), filter).Decode(&p); err != nil {
		log.Fatal(err)
	}
	return p
}

func GetAll() []Dataa {
	collec := db.GetCollection("maindatas")
	filter := bson.M{}
	var p1 []Dataa
	cursor, err := collec.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var p Dataa
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Fatal(err)
		}
		p1 = append(p1, p)
	}
	return p1
}

func Update(alias string, new Dataa) int64 {
	collec := db.GetCollection("maindatas")
	update := bson.M{"$set": bson.M{"updatedate": time.Now()}}
	filter := bson.M{"fulllink": alias}
	res, err := collec.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return res.ModifiedCount
}

func Remove(alias string) int64 {
	collec := db.GetCollection("maindatas")
	filter := bson.M{"$or": []bson.M{bson.M{"fulllink": alias}, bson.M{"shortenlink": alias}}}
	res, err := collec.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return res.DeletedCount
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
