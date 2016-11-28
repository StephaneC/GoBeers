package main

import (
  "fmt"
	"log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Beer struct {
        Name string
        Alcohol float32
}

func addBeer(b Beer){
  session, err := mgo.Dial("server1.example.com,server2.example.com")
          if err != nil {
                  panic(err)
          }
          defer session.Close()

          // Optional. Switch the session to a monotonic behavior.
          session.SetMode(mgo.Monotonic, true)

          //choose collection from db
          c := session.DB("test").C("beers")
          err = c.Insert(b)
          if err != nil {
                  log.Fatal(err)
          }

          result := Person{}
          err = c.Find(bson.M{"name": "Ale"}).One(&result)
          if err != nil {
                  log.Fatal(err)
          }
}

func getBeers() *Beer{
  session, err := mgo.Dial("server1.example.com,server2.example.com")
          if err != nil {
                  panic(err)
          }
          defer session.Close()

          // Optional. Switch the session to a monotonic behavior.
          session.SetMode(mgo.Monotonic, true)

          //choose collection from db
          c := session.DB("test").C("beers")
          result := Beer{}
          err = c.Find(bson.M{}).One(&result)
          if err != nil {
                  log.Fatal(err)
          }
}
