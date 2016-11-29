package main

import (
	"log"
	"os"
	"encoding/json"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

const collection = "beers"
const dbName = "BrestOpenCampus"

type Beer struct {
        Name string
        Alcohol float32
}

var (
	mgoSession   *mgo.Session
	databaseName = "myDB"
	err          error
)

func openDbConnection() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("127.0.0.1")//os.Getenv("dbHost"))
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func addBeer(b Beer){
  openDbConnection();

  //choose collection from db
  c := mgoSession.DB(dbName).C(collection)
  err = c.Insert(b)
  if err != nil {
    log.Fatal(err)
  }
}

func getBeers() string{
  openDbConnection();

  //choose collection from db
  c := mgoSession.DB(dbName).C(collection)
  var results []Beer
  err = c.Find(bson.M{}).All(&results)
  if err != nil {
    log.Fatal(err)
  }
  jsonBeer, _ := json.Marshal(results)
  return string(jsonBeer)
}
