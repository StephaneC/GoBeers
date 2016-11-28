package main

import (
	"log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"  //for var env
)

const collection = "beers"
const dbName = "BrestOpenCampus"

var dbHost = os.GetEnv("dbHost")
var dbPwd = os.GetEnv("dbPwd")

type Beer struct {
        Name string
        Alcohol float32
}

func openDbConnection() *mgo.Session {
  session, err := mgo.Dial(dbHost)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  return session
}

func addBeer(b Beer){
  openDbConnection();

  //choose collection from db
  c := session.DB(dbName).C(collection)
  err = c.Insert(b)
  if err != nil {
    log.Fatal(err)
  }

  result := Person{}
  err = c.Find(bson.M{}).One(&result)
  if err != nil {
    log.Fatal(err)
  }
}

func getBeers() string{
  openDbConnection();

  //choose collection from db
  c := session.DB(dbName).C(collection)
  result := []Beer
  err = c.Find(bson.M{}).One(&result)
  if err != nil {
    log.Fatal(err)
  }
  jsonBeer, _ := json.Marshal(result)
  return string(jsonBeer)
}
