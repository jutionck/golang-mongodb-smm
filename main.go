package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const mongodbUri = "mongodb://207.148.125.99:27017"

type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"fullName"`
	Age      int                `bson:"age"`
	Gender   string             `bson:"gender"`
	JoinDate primitive.DateTime `bson:"joinDate"`
	Senior   bool               `bson:"senior"`
}

func main() {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "jutionck",
		Password:      "password",
	}
	clientOptions := options.Client()
	clientOptions.ApplyURI(mongodbUri).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connect, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := connect.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// membuat sebuah db - collection
	db := connect.Database("enigma")
	coll := db.Collection("student")

	// Create insertOne
	//newId, err := coll.InsertOne(ctx, bson.D{
	//	{"age", 19},
	//	{"name", "Bulan"},
	//	{"gender", "F"},
	//	{"senior", false},
	//})

	// slice
	//jd01 := parseTime("2022-07-02 15:04:05")
	//jd02 := parseTime("2022-07-03 15:04:05")
	////jd03 := parseTime("2022-07-04 15:04:05")
	//students := []interface{}{
	//	bson.D{
	//		{"name", "Sita"},
	//		{"age", 29},
	//		{"gender", "F"},
	//		{"joinDate", primitive.NewDateTimeFromTime(jd01)},
	//		{"senior", true},
	//	},
	//	bson.D{
	//		{"name", "Melani"},
	//		{"age", 25},
	//		{"gender", "F"},
	//		{"joinDate", jd02},
	//		{"senior", true},
	//	},
	//	bson.D{
	//		{"name", "Suci"},
	//		{"age", 10},
	//		{"gender", "F"},
	//		{"joinDate", primitive.NewDateTimeFromTime(parseTime("2022-07-13"))},
	//		{"senior", false},
	//	},
	//}

	//newStudent := Student{
	//	Id:       primitive.NewObjectID(),
	//	Name:     "Dino",
	//	Age:      21,
	//	Gender:   "M",
	//	JoinDate: primitive.NewDateTimeFromTime(parseTime("2022-07-13 00:00:00")),
	//	Senior:   false,
	//}
	//
	//newId, err := coll.InsertOne(ctx, newStudent)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Printf("inserted document with ID %v\n", newId.InsertedID)

	// DeleteOne
	// DeleteOne(ctx, filter)
	//del01, err := coll.DeleteOne(ctx, bson.D{{"age", bson.D{{"$lt", 20}}}})
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Printf("Number of documents deleted: %d\n", del01.DeletedCount)
	//
	//// UpdateOne
	//// UpdateOne(ctx, filter, updateValue)
	//upd01, err := coll.UpdateOne(ctx,
	//	bson.D{{"fullName", "Dona"}},
	//	bson.D{{"$set", bson.D{{"joinDate", primitive.NewDateTimeFromTime(parseTime("2022-07-13 00:00:00"))}}}})
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Printf("Documents matched: %v\n", upd01.MatchedCount)
	//fmt.Printf("Documents updated: %v\n", upd01.ModifiedCount)

	// Read
	// SELECT * FROM student
	cursor, err := coll.Find(ctx, bson.D{}, options.Find().SetProjection(bson.D{
		{"_id", 1},
		{"name", 1},
		{"gender", 1},
	}))
	if err != nil {
		log.Println(err.Error())
	}
	var students []bson.D
	err = cursor.All(ctx, &students)
	if err != nil {
		log.Println(err.Error())
	}
	for _, student := range students {
		fmt.Println("SELECT ALL: ", student)
	}

	// Logical
	filterGenderAndAge := bson.D{
		{"$and", bson.A{
			bson.D{
				{"gender", "M"},
				{"age", bson.D{{"$gte", 10}}},
			},
		}},
	}
	projection := bson.D{
		{"_id", 1},
		{"fullName", 1},
		{"gender", 1},
		{"age", 1},
	}
	//cursor, err = coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = cursor.All(ctx, &students)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// mapping result query ke struct
	filterGenderAndAgeResult := make([]*Student, 0)
	cursor, err = coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))
	if err != nil {
		log.Println(err.Error())
	}
	for cursor.Next(ctx) {
		var student Student
		err := cursor.Decode(&student)
		if err != nil {
			log.Println(err.Error())
		}
		filterGenderAndAgeResult = append(filterGenderAndAgeResult, &student)
	}
	for _, student := range filterGenderAndAgeResult {
		fmt.Println("FILTER BY GENDER & AGE (WITH STRUCT)", student)
	}

	// Aggregation
	coll = connect.Database("enigma").Collection("products")
	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Product total: ", count)

	// with filter
	count, err = coll.CountDocuments(ctx, bson.D{{"category", "food"}})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Product total with category[food]: ", count)

	// Match, group, sort, dll
	matchStage := bson.D{
		{"$match", bson.D{
			{"category", "food"},
		}},
	}

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$category"},
			{"Total", bson.D{{"$sum", 1}}},
		}},
	}
	cursor, err = coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		log.Println(err.Error())
	}
	var productCount []bson.M
	err = cursor.All(ctx, &productCount)
	if err != nil {
		log.Println(err.Error())
	}
	for _, product := range productCount {
		fmt.Printf("Group[%v], Total[%v]\n ", product["_id"], product["Total"])
	}

}

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}

/**
* Buat koneksi ke mongodb (url) -> mongodb://localhost:27017
* Siapkan User Auth: username & password
 */
